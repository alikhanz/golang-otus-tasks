package main

import (
	"errors"
	"fmt"
	"github.com/alikhanz/golang-otus-tasks/hw-5-task-runner/runner"
)

func main() {
	r := runner.New()

	tasks := make([]func()error, 0, 1)

	tasks = append(tasks, func() error {
		return errors.New("Mew mew")
	})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	//tasks = append(tasks, func() error {
	//	return nil
	//})
	err := r.Run(tasks, 1, 1)

	fmt.Println(err, r)
}