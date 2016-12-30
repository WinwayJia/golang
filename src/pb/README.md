###
loginServer: 登录，分配msgServer给client，接受msgServer的负载报告
RouteServer: 转发位于不同msgServer的client的消息
msgServer: client带着loginServer分配的token登录，接受用户消息，给用户发送消息

msgServer通过redis队列发送负载报告给loginServer

