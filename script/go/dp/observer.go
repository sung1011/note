package main

// 报社 —— 客户
type Customer interface {
	update()
}

type CustomerA struct {
	recv bool
}

func (p *CustomerA) update() {
	p.recv = true
	// fmt.Println("客户A收到报纸")
}

type CustomerB struct {
	recv bool
}

func (p *CustomerB) update() {
	p.recv = true
	// fmt.Println("客户B收到报纸")
}

// 报社 （被观察者)
type NewsOffice struct {
	customers []Customer
}

func (n *NewsOffice) addCustomer(customer Customer) {
	n.customers = append(n.customers, customer)
}

func (n *NewsOffice) newspaperCome() {
	// 主体逻辑...

	// 通知所有客户
	n.notifyAllCustomer()
}

func (n *NewsOffice) notifyAllCustomer() {
	for _, customer := range n.customers {
		customer.update()
	}
}
