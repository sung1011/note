# MongoDB 副本命令

## Replication Methods in the mongo Shell

| Name                           | Description                                                                                                                    |
| ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| rs.add()                       | Adds a member to a replica set. 将成员添加到副本集                                                                             |
| rs.addArb()                    | Adds an  to a replica set. 将仲裁节点添加到副本集                                                                              |
| rs.conf()                      | Returns the replica set configuration document. 返回副本集的配置内容                                                           |
| rs.freeze()                    | Prevents the current member from seeking election as primary for a period of time. 阻止当前成员在一段时间内寻求选举为主节点    |
| rs.help()                      | Returns basic help text for  functions. 返回功能的基本帮助文本                                                                 |
| rs.initiate()                  | Initializes a new replica set. 初始化新的副本集                                                                                |
| rs.printReplicationInfo()      | Prints a report of the status of the replica set from the perspective of the primary. 以主节点的角度来打印副本集状态的报告     |
| rs.printSlaveReplicationInfo() | Prints a report of the status of the replica set from the perspective of the secondaries. 以从节点的角度来打印副本集状态的报告 |
| rs.reconfig()                  | Re-configures a replica set by applying a new replica set configuration object. 通过应用新的副本集配置对象来重新配置副本集     |
| rs.remove()                    | Remove a member from a replica set. 将成员从副本集中移除                                                                       |
| rs.status()                    | Returns a document with information about the state of the replica set. 返回包含关于副本集状态信息的文档                       |
| rs.stepDown()                  | Causes the current  to become a secondary which forces an . 使当前的转变为从节点，同时触发                                     |
| rs.syncFrom()                  | 设置复制集成员从哪个成员中同步数据，同时覆盖默认的同步目标选择逻辑。                                                           |

## Replication Database Commands

| Name                       | Description                                                                                                                        |
| -------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| applyOps                   | Internal command that applies  entries to the current data set. 内部命令，可将条目应用于当前数据集                                 |
| isMaster                   | 显示关于此成员在副本集中的角色信息，包括它是否为主角色                                                                             |
| replSetAbortPrimaryCatchUp | 对所选的主节点强行中止同步（即追平数据），然后完成到主节点的转换                                                                   |
| replSetFreeze              | Prevents the current member from seeking election as  for a period of time. 阻止当前成员在一段时间内寻求选举为                     |
| replSetGetConfig           | Returns the replica set’s configuration object. 返回副本集的配置对象                                                               |
| replSetGetStatus           | Returns a document that reports on the status of the replica set. 返回报告副本集状态的文档                                         |
| replSetInitiate            | Initializes a new replica set. 初始化新的副本集                                                                                    |
| replSetMaintenance         | Enables or disables a maintenance mode, which puts a  node in a RECOVERING state. 启用或禁用维护模式，该模式会将置于RECOVERING状态 |
| replSetReconfig            | Applies a new configuration to an existing replica set. 将新的配置应用于现有副本集                                                 |
| replSetResizeOplog         | 动态调整副本集成员oplog的大小。该功能仅适用于WiredTiger存储引擎                                                                    |
| replSetStepDown            | Forces the current  to step down and become a , forcing an election. 使当前的转变为,，同时触发                                     |
| replSetSyncFrom            | Explicitly override the default logic for selecting a member to replicate from. 显式重写用于选择要复制的成员的默认逻辑。           |