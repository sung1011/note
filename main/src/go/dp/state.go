package main

// 这是一个工作流的例子, 在企业内部或者是学校我们经常会看到很多审批流程
// 假设我们有一个报销的流程: 员工提交报销申请 -> 直属部门领导审批 -> 财务审批 -> 结束
// 在这个审批流中, 处在不同的环节就是不同的状态
// 而流程的审批、驳回就是不同的事件

import "fmt"

// Machine 状态机
type Machine struct {
	state IState
}

// SetState 更新状态
func (m *Machine) SetState(state IState) {
	m.state = state
}

func (m *Machine) GetStateName() string {
	return m.state.GetName()
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

// IState 状态
type IState interface {
	// 审批通过
	Approval(m *Machine)
	// 驳回
	Reject(m *Machine)
	// 获取当前状态名称
	GetName() string
}

type leaderApproveState struct{}

func (leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader 审批成功 -> 进入财务流程")
	m.SetState(GetFinanceApproveState()) // 修改对象主体, 即改变了其后续行为
}

func (leaderApproveState) GetName() string {
	return "LeaderState"
}

func (leaderApproveState) Reject(m *Machine) {}

func GetLeaderApproveState() IState {
	return &leaderApproveState{}
}

type financeApproveState struct{}

func (f financeApproveState) Approval(m *Machine) {
	fmt.Println("财务审批成功 打款")
}

func (f financeApproveState) Reject(m *Machine) {
	fmt.Println("财务拒绝 -> 回到leader审批")
	m.SetState(GetLeaderApproveState())
}

func (f financeApproveState) GetName() string {
	return "FinanceState"
}

func GetFinanceApproveState() IState {
	return &financeApproveState{}
}
