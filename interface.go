package main

import (
	"gitlab.szzhijing.com/quanbu/edge/park-edge-qms/model"
)

type EngineInterface interface {

	/*
	  引擎启动（应用程序启动时只调用一次），需指定算法模块模型路径
	*/
	Open(modelFile string) error
	/*
	  引擎关闭（应用程序关闭时只调用一次）
	*/
	Close() error
	/*
	  引擎初始(初始服务对象后调用检查接口前需调用，以及引擎重置后需重新进入检查状态前调用，详见调用流程图)
	*/
	Init(req model.InitReqDTO) error
	/*
	  引擎重置(销毁服务对象前调用落布接口后需调用，以及落布调用后需重新进入检查状态前调用，详见调用流程图)
	*/
	Reset() error

	/*
	  检查接口
	*/
	Detect(req model.DetectReqDTO) ([]model.DefectInfo, error)
	/*
	  无图拼接（测试接口）
	*/
	Split(req model.DetectReqDTO1) ([]model.DefectInfo, error)
	/*
	  落布接口
	*/
	Shutdown() ([]model.DefectInfo, error)
	/*
	  织机工单事件
	*/
	LoomEvent(req model.LoomReqDTO) (model.M, error)

	/*
	  获取sdk相关版本接口信息
	*/
	GetVersion() (model.VersionInfo, error)
}
