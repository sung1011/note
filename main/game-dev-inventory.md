# game dev inventory

## id 唯一id

- 自增
- uuid

## template_id 模板id

确定模板的基础属性

```js
{
    template_id: "weapon_ak47",
    name: "AK-47",
    type: "weapon",
    rarity: "epic",
    attributes: {
        damage: 35,
        range: 500,
        accuracy: 0.8
    }
}
```

## 词条

- g位码 高效查询词条是否存在, 节约空间

```go
const (
    MASK_DAMAGE   = 1 << iota
    MASK_RANGE
    MASK_ACCURACY
)

type Item struct {
    ID         string
    TemplateID string
    Name       string
    Type       string
    Rarity     string
    Attributes map[string]interface{}
    GBitmask   uint64 // 位图存储词条
}

func (i *Item) HasAttribute(attr string) bool {
    switch attr {
    case "damage":
        return i.GBitmask&MASK_DAMAGE != 0
    case "range":
        return i.GBitmask&MASK_RANGE != 0
    case "accuracy":
        return i.GBitmask&MASK_ACCURACY != 0
    default:
        return false
    }
}
```

## 背包整理

- 预测,回滚 关闭背包时才上传服务器, 服务器进行校验



