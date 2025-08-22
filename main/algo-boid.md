# 鸟群算法

**鸟群算法（Boids Algorithm）** 是由计算机图形学专家 Craig Reynolds 在 1987 年提出的一种模拟群体行为的算法。该算法通过简单的局部规则，能够产生复杂而自然的群体运动模式，广泛应用于游戏开发、动画制作、机器人群体控制等领域。

---

## **算法原理**

鸟群算法的核心思想是：**每个个体（boid）仅根据其邻近个体的状态，遵循三个简单规则进行决策**，从而在宏观上形成协调一致的群体行为。

这种"**局部交互产生全局模式**"的现象被称为**涌现行为（Emergent Behavior）**。

---

## **三大核心规则**

### 1. **分离（Separation）**

- **目标**：避免与邻近个体过于拥挤
- **行为**：当周围个体距离过近时，产生远离的力
- **作用**：防止碰撞，保持个体间的最小安全距离

### 2. **对齐（Alignment）**

- **目标**：与邻近个体保持相同的运动方向
- **行为**：调整自身速度方向，趋向于邻居的平均速度方向
- **作用**：使群体运动方向趋于一致

### 3. **聚合（Cohesion）**

- **目标**：向邻近个体的中心位置靠拢
- **行为**：产生朝向邻居平均位置的吸引力
- **作用**：保持群体的凝聚性，避免个体脱离群体

---

## **算法实现**

### 基本数据结构

```javascript
class Boid {
    constructor(x, y) {
        this.position = { x, y };
        this.velocity = { x: 0, y: 0 };
        this.acceleration = { x: 0, y: 0 };
        this.maxSpeed = 2;
        this.maxForce = 0.03;
        this.perceptionRadius = 50; // 感知半径
    }
}
```

### 核心更新逻辑

```javascript
function updateBoid(boid, flock) {
    // 计算三个力的合力
    const sep = separate(boid, flock);
    const ali = align(boid, flock);
    const coh = cohesion(boid, flock);
    
    // 应用权重
    sep.multiply(1.5);  // 分离力权重
    ali.multiply(1.0);  // 对齐力权重
    coh.multiply(1.0);  // 聚合力权重
    
    // 合成加速度
    boid.acceleration.add(sep);
    boid.acceleration.add(ali);
    boid.acceleration.add(coh);
    
    // 更新速度和位置
    boid.velocity.add(boid.acceleration);
    boid.velocity.limit(boid.maxSpeed);
    boid.position.add(boid.velocity);
    boid.acceleration.multiply(0); // 重置加速度
}
```

### 分离规则实现

```javascript
function separate(boid, flock) {
    const desiredSeparation = 25;
    let steer = { x: 0, y: 0 };
    let count = 0;
    
    for (let other of flock) {
        const distance = getDistance(boid.position, other.position);
        
        if (distance > 0 && distance < desiredSeparation) {
            // 计算远离方向
            const diff = subtract(boid.position, other.position);
            normalize(diff);
            diff.divide(distance); // 距离越近，力越大
            steer.add(diff);
            count++;
        }
    }
    
    if (count > 0) {
        steer.divide(count);
        normalize(steer);
        steer.multiply(boid.maxSpeed);
        steer.subtract(boid.velocity);
        steer.limit(boid.maxForce);
    }
    
    return steer;
}
```

### 对齐规则实现

```javascript
function align(boid, flock) {
    let sum = { x: 0, y: 0 };
    let count = 0;
    
    for (let other of flock) {
        const distance = getDistance(boid.position, other.position);
        
        if (distance > 0 && distance < boid.perceptionRadius) {
            sum.add(other.velocity);
            count++;
        }
    }
    
    if (count > 0) {
        sum.divide(count);
        normalize(sum);
        sum.multiply(boid.maxSpeed);
        
        const steer = subtract(sum, boid.velocity);
        steer.limit(boid.maxForce);
        return steer;
    }
    
    return { x: 0, y: 0 };
}
```

### 聚合规则实现

```javascript
function cohesion(boid, flock) {
    let sum = { x: 0, y: 0 };
    let count = 0;
    
    for (let other of flock) {
        const distance = getDistance(boid.position, other.position);
        
        if (distance > 0 && distance < boid.perceptionRadius) {
            sum.add(other.position);
            count++;
        }
    }
    
    if (count > 0) {
        sum.divide(count);
        return seek(boid, sum); // 朝向目标位置
    }
    
    return { x: 0, y: 0 };
}

function seek(boid, target) {
    const desired = subtract(target, boid.position);
    normalize(desired);
    desired.multiply(boid.maxSpeed);
    
    const steer = subtract(desired, boid.velocity);
    steer.limit(boid.maxForce);
    return steer;
}
```

---

## **扩展功能**

### 1. **避障行为**

```javascript
function avoid(boid, obstacles) {
    // 检测前方障碍物并产生避让力
    const ahead = boid.position.copy();
    ahead.add(boid.velocity.copy().normalize().multiply(50));
    
    for (let obstacle of obstacles) {
        if (getDistance(ahead, obstacle.position) < obstacle.radius) {
            const avoidance = subtract(ahead, obstacle.position);
            normalize(avoidance);
            avoidance.multiply(boid.maxSpeed);
            return avoidance;
        }
    }
    
    return { x: 0, y: 0 };
}
```

### 2. **领导者跟随**

```javascript
function follow(boid, leader) {
    const behind = leader.position.copy();
    behind.subtract(leader.velocity.copy().normalize().multiply(30));
    return seek(boid, behind);
}
```

### 3. **边界处理**

```javascript
function wrapAround(boid, width, height) {
    if (boid.position.x < 0) boid.position.x = width;
    if (boid.position.x > width) boid.position.x = 0;
    if (boid.position.y < 0) boid.position.y = height;
    if (boid.position.y > height) boid.position.y = 0;
}
```

---

## **性能优化**

### 1. **空间分割**

使用四叉树或网格划分减少邻居搜索的计算复杂度：

```javascript
// 将空间划分为网格，只检查相邻网格中的个体
const grid = new SpatialGrid(cellSize);
grid.insert(boid);
const neighbors = grid.getNearby(boid.position, boid.perceptionRadius);
```

### 2. **距离计算优化**

使用平方距离避免开方运算：

```javascript
function getSquaredDistance(p1, p2) {
    const dx = p1.x - p2.x;
    const dy = p1.y - p2.y;
    return dx * dx + dy * dy;
}
```

---

## **应用场景**

### 1. **游戏开发**

- **NPC 群体行为**：士兵编队、动物群体
- **粒子系统**：鱼群、鸟群效果
- **敌人 AI**：群体围攻、协同作战

### 2. **动画制作**

- **电影特效**：大规模群体场景
- **程序化动画**：自然的群体运动

### 3. **机器人学**

- **无人机编队**：协调飞行
- **机器人群体**：协同探索、搬运

### 4. **交通模拟**

- **车辆流动**：交通流建模
- **行人模拟**：人群疏散

---

## **参数调优技巧**

1. **权重平衡**：分离 > 对齐 ≈ 聚合，避免过于拥挤
2. **感知半径**：过小导致群体分散，过大影响性能
3. **最大速度**：影响群体的活跃程度
4. **最大转向力**：控制个体的机动性

---

## **总结**

鸟群算法展现了**简单规则产生复杂行为**的经典案例。通过三个基础规则的组合，能够模拟出自然界中各种群体行为。其优势在于：

- **去中心化**：无需全局控制器
- **可扩展性**：易于添加新的行为规则
- **自然性**：产生真实的群体运动模式
- **实时性**：计算效率高，适合实时应用

该算法不仅在计算机图形学中具有重要地位，也为人工智能、复杂系统研究提供了宝贵的启发。

