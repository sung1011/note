# delta time

在游戏前端技术中，**Delta Time（ΔT 或 deltaTime）** 是指**当前帧与上一帧之间的时间差**（通常以秒为单位），它在游戏开发中扮演着关键角色，主要用于解决不同设备帧率（FPS）不一致带来的问题，确保游戏逻辑和动画的稳定运行。

---

## **Delta Time 的作用**

1. **帧率无关的平滑动画**  
   如果直接基于帧数更新物体位置（例如：`物体位置 += 速度`），高帧率设备（如 120 FPS）的物体会移动得更快，而低帧率设备（如 30 FPS）会更慢。Delta Time 通过将速度与时间差绑定（例如：`物体位置 += 速度 * deltaTime`），保证物体在任何帧率下每秒移动的距离一致。

2. **物理模拟的稳定性**  
   物理计算（如碰撞检测、重力加速度）对时间敏感。直接使用帧率可能导致计算错误，比如：

   ```javascript
   // 错误示例（帧率低时跳跃过大，可能“穿墙”）
   positionY += gravity * framesPassed;
   ```

   Delta Time 可以逐帧细化计算，避免跳跃：

   ```javascript
   // 正确示例（每秒速度均匀变化）
   positionY += gravity * deltaTime;
   ```

3. **性能波动时的容错性**  
   当设备卡顿时（如某一帧耗时 0.5 秒），Delta Time 的存在会告诉游戏：“这帧花了较多时间，需要按比例调整逻辑”，避免角色瞬移或动画骤停。

---

## **实现方式**

在游戏循环中，每一帧记录当前时间并计算与上一帧的时间差：

```javascript
let lastTime = Date.now();

function gameLoop() {
    const currentTime = Date.now();
    const deltaTime = (currentTime - lastTime) / 1000; // 转为秒
    lastTime = currentTime;

    updateGameLogic(deltaTime); // 将 deltaTime 传递给逻辑更新函数
    render();
    requestAnimationFrame(gameLoop);
}
```

在物体更新时，乘上 `deltaTime`：

```javascript
function updateGameLogic(deltaTime) {
    player.x += movementSpeed * deltaTime; // 保持每秒移动速度固定
    enemy.position = lerp(start, end, t * deltaTime); // 平滑插值
}
```

---

## **需要注意的问题**

1. **处理极端卡顿**  
   当 `deltaTime` 过大时（如超过 0.1 秒），可能导致物理模拟错误（物体穿过墙壁）。常见解决方案：
   - **限制最大 Delta Time**：例如每帧最多允许处理 0.1 秒的间隔。
   - **分步更新**：将一个大 `deltaTime` 拆分为多个小步长迭代计算（如每次处理 0.016 秒）。

2. **固定时间步长（Fixed Timestep）**  
   物理引擎（如 Box2D）通常使用固定的时间步长（例如 1/60 秒）来保证计算稳定性，同时允许渲染帧率独立。

---

## **实际应用场景**

- **角色移动、动画插值**：确保角色在不同设备上的移动速度一致。
- **粒子效果**：粒子运动和消亡时间需要与真实时间同步。
- **联机游戏**：网络同步逻辑可能需要基于时间差调整状态。

---

## **总结**

Delta Time 是游戏开发的**核心基础概念**，它的本质是将游戏逻辑的更新与时间流逝解耦，使游戏的表现不受硬件性能或帧率波动的影响，从而提供一致的体验。无论是 Unity 的 `Time.deltaTime` 还是 Unreal 的 `DeltaSeconds`，其原理都与此一致。