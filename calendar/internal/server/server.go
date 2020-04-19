package server

import (
	"context"
	"fmt"
	"github.com/alikhanz/golang-otus-tasks/calendar/internal/mapper"
	"github.com/alikhanz/golang-otus-tasks/calendar/internal/middlewares"
	"github.com/alikhanz/golang-otus-tasks/calendar/pkg/calendar"
	"github.com/alikhanz/golang-otus-tasks/calendar/pkg/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/urfave/negroni"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
)

type Config struct {
	GrpcPort   int
	HttpListen string
}

type Server struct {
	Config    Config
	calServer *CalendarServer
}

func NewServer(config Config, cal *calendar.Calendar) *Server {
	s := &Server{Config: config}
	s.calServer = NewCalendarServer(cal)

	return s
}

func (s *Server) Run() error {
	eg := errgroup.Group{}
	eg.Go(s.grpcInit)
	eg.Go(s.httpInit)

	return eg.Wait()
}

func (s *Server) grpcInit() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Config.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalendarServer(grpcServer, s.calServer)
	return grpcServer.Serve(lis)
}

func (s *Server) httpInit() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	n := negroni.New(negroni.NewRecovery(), middlewares.NewLogger())
	n.UseHandler(mux)

	err := pb.RegisterCalendarHandlerServer(ctx, mux, s.calServer)
	if err != nil {
		return err
	}
	return http.ListenAndServe(s.Config.HttpListen, n)
}

var eventMapper *mapper.EventMapper

type CalendarServer struct {
	cal       *calendar.Calendar
	validator *Validator
}

func NewCalendarServer(cal *calendar.Calendar) *CalendarServer {
	v := NewValidator(cal)
	return &CalendarServer{cal: cal, validator: v}
}

func (c CalendarServer) CreateEvent(
	ctx context.Context,
	request *pb.CreateEventRequest,
) (*pb.Event, error) {
	err := c.validator.ValidateCreateEventRequest(request)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	t, err := ptypes.Timestamp(request.Time)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid date")
	}

	eId, err := c.cal.CreateEvent(
		request.Title,
		request.Description,
		t,
		request.Repeatable,
	)

	if err != nil {
		return nil, status.Error(codes.Internal, "Cannot save event")
	}

	ev, err := c.cal.FetchById(eId)

	if err != nil {
		return nil, status.Error(codes.Internal, "Event fetch failed")
	}

	return eventMapper.MapModelToPb(ev), nil
}

func (c CalendarServer) UpdateEvent(
	ctx context.Context,
	event *pb.Event,
) (*pb.Event, error) {
	err := c.validator.ValidateEventFields(event)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = c.validator.ValidateEventExist(event)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ev, err := eventMapper.MapPbToModel(event)

	if err != nil {
		// Если это произошло, значит что-то не довалидировали перед маппингом эвента.
		return nil, status.Error(codes.Internal, "Internal error")
	}

	err = c.cal.UpdateEvent(ev)

	if err != nil {
		return nil, status.Error(codes.Internal, "Cannot save event")
	}

	return eventMapper.MapModelToPb(ev), nil
}

func (c CalendarServer) RemoveEvent(
	ctx context.Context,
	request *pb.RemoveEventRequest,
) (*pb.RemoveEventResponse, error) {
	id, err := uuid.Parse(request.EventId)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Incorrect uuid")
	}

	err = c.cal.RemoveEvent(id)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Cannot find event")
	}

	return &pb.RemoveEventResponse{}, nil
}

func (c CalendarServer) GetEventsList(
	ctx context.Context,
	request *pb.GetEventsListRequest,
) (*pb.EventsList, error) {
	ft, err := ptypes.Timestamp(request.FromDate)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "FromDate incorrect")
	}

	tt, err := ptypes.Timestamp(request.ToDate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "ToDate incorrect")
	}

	result, err := c.cal.FetchBetweenDates(ft, tt)

	return &pb.EventsList{Events: eventMapper.MapModelListToPb(result)}, nil
}
