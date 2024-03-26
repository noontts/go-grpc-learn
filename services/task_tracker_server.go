package services

import (
  context "context"
)

type tasksTrackerServer struct{
}

func NewTasksTrackerServer() TasksTrackerServer {
  return tasksTrackerServer{}
} 

func (tasksTrackerServer) mustEmbedUnimplementedTasksTrackerServer(){}


func (tasksTrackerServer) CreateNewTasks(ctx context.Context, req *CreateTasksRequest) (*CreateTasksResponse, error){}

func (tasksTrackerServer) GetListTasks(context.Context, *GetListTasksRequest) (*GetListTasksResponse, error){}

  


