
// 用户信息相关函数
function refreshUserInfo() {
    try {
        if (typeof getUserInfo !== 'function') {
            showToast('接口未加载', true);
            return;
        }

        const userInfoStr = getUserInfo();
        const userInfo = JSON.parse(userInfoStr);

        const userEquipAttributesStr = getUserEquipAttributes();
        const userEquipAttributes = JSON.parse(userEquipAttributesStr);

        document.getElementById('username').textContent = userInfo.username || '未设置';
        document.getElementById('cultivation').textContent = userInfo.cultivation ;

        let hpStr = userInfo.hp + '/' + userInfo.hpLimit;
        if(userEquipAttributes.Hp!==0){hpStr += '(+' + userEquipAttributes.Hp+')';}
        document.getElementById('hp').textContent = hpStr;
        let attackStr = userInfo.attack;
        if(userEquipAttributes.Attack!==0){attackStr += '(+' + userEquipAttributes.Attack+')';}
        document.getElementById('attack').textContent =attackStr;
        let defenseStr = userInfo.defense;
        if(userEquipAttributes.Defense!==0){defenseStr += '(+' + userEquipAttributes.Defense+')';}
        document.getElementById('defense').textContent = defenseStr;
        let speedStr = userInfo.speed;
        if(userEquipAttributes.Speed!==0){speedStr += '(+' + userEquipAttributes.Speed+')';}
        document.getElementById('speed').textContent = speedStr;

        document.getElementById('exp').textContent = userInfo.exp  + '/' + userInfo.level*10; // 转为数字
        document.getElementById('potential').textContent = userInfo.potential || '0';
        document.getElementById('gold').textContent = userInfo.gold || '0';

        const btn = document.getElementById('restartB');
        if (btn) {
            if(userInfo.level>=150){
                btn.classList.remove('hidden');
            }else{
                btn.classList.add('hidden'); // 添加隐藏类
            }
        }

        showToast('信息已刷新');
    } catch (e) {
        showToast('获取信息失败: ' + e.message, true);
    }
}

async function setUsernameBtn() {
    const newUsername = document.getElementById('newUsername').value.trim();
    if (!newUsername) {
        showToast('请输入用户名', true);
        return;
    }

    try {
        if (typeof setUsername !== 'function') {
            showToast('接口未加载', true);
            return;
        }

        setUsername(newUsername);
        document.getElementById('newUsername').value = '';
        refreshUserInfo();
        showToast('用户名已更新');
        addToLog("更名成功!")
    } catch (e) {
        showToast('设置失败: ' + e.message, true);
    }
}

// 分配属性点
async function allocateBtn(stat) {
    try {
        if (typeof allocate !== 'function') {
            showToast('接口未加载', true);
            return;
        }

        const result = allocate(stat);
        refreshUserInfo();
        showToast(result);
        addToLog(`分配潜能点: ${result}`);
    } catch (e) {
        showToast('操作失败: ' + e.message, true);
    }
}

// 全局变量管理定时器状态
let activeTimer = null;       // 当前活跃的定时器
let currentAction = null;     // 当前执行的动作类型

// 辅助函数：只禁用其他按钮，保留当前活动按钮可点击
function disableOtherButtons(activeBtnId) {
    // 先启用所有按钮，避免状态混乱
    enableAllButtons();

    // 禁用除当前活动按钮外的其他按钮
    document.querySelectorAll('#healBtn, #cultivationBtn, #getGoldBtn').forEach(btn => {
        if (btn.id !== activeBtnId) {
            btn.disabled = true;
            btn.classList.add('opacity-50', 'cursor-not-allowed');
        }
    });
}

// 辅助函数：启用所有按钮
function enableAllButtons() {
    document.querySelectorAll('#healBtn, #cultivationBtn, #getGoldBtn').forEach(btn => {
        btn.disabled = false;
        btn.classList.remove('opacity-50', 'cursor-not-allowed');
    });
}

// 停止当前循环（完整实现）
function stopCurrentLoop() {
    if (activeTimer) {
        clearInterval(activeTimer);  // 清除定时器
        activeTimer = null;          // 重置定时器变量
        currentAction = null;        // 重置当前动作
        enableAllButtons();          // 启用所有按钮
        addToLog('已停止循环操作');
    }
}

// 恢复健康（循环版）
async function healBtn() {
    // 如果已有活跃循环且是当前动作，则停止
    if (activeTimer && currentAction === 'heal') {
        stopCurrentLoop();
        return;
    }

    // 如果有其他活跃循环，先停止
    if (activeTimer) {
        stopCurrentLoop();
    }

    // 标记当前动作并禁用其他按钮
    currentAction = 'heal';
    disableOtherButtons('healBtn'); // 传入当前按钮ID

    // 立即执行一次，然后每分钟循环
    await executeHeal();
    activeTimer = setInterval(executeHeal, 60 * 1000); // 60秒间隔
    addToLog('已开始自动恢复，每分钟一次（点击按钮停止），`修炼`、`做工`、`恢复`同一时间段只能进行其中一个操作');
}

// 修炼（循环版）
async function cultivationBtn() {
    if (activeTimer && currentAction === 'cultivation') {
        stopCurrentLoop();
        return;
    }

    if (activeTimer) {
        stopCurrentLoop();
    }

    currentAction = 'cultivation';
    disableOtherButtons('cultivationBtn'); // 传入当前按钮ID

    await executeCultivation();
    activeTimer = setInterval(executeCultivation, 60 * 1000);
    addToLog('已开始自动修炼，每分钟一次（点击按钮停止），`修炼`、`做工`、`恢复`同一时间段只能进行其中一个操作');
}

// 做工赚钱（循环版）
async function getGoldBtn() {
    if (activeTimer && currentAction === 'getGold') {
        stopCurrentLoop();
        return;
    }

    if (activeTimer) {
        stopCurrentLoop();
    }

    currentAction = 'getGold';
    disableOtherButtons('getGoldBtn'); // 传入当前按钮ID

    await executeGetGold();
    activeTimer = setInterval(executeGetGold, 60 * 1000);
    addToLog('已开始自动做工，每分钟一次（点击按钮停止），`修炼`、`做工`、`恢复`同一时间段只能进行其中一个操作');
}

// 实际执行恢复的函数
async function executeHeal() {
    try {
        if (typeof heal !== 'function') {
            showToast('恢复接口未加载', true);
            stopCurrentLoop(); // 接口异常时停止循环
            return;
        }

        const result = heal();
        refreshUserInfo();
        showToast(result);
        addToLog(`恢复健康: ${result}`);
    } catch (e) {
        showToast('恢复失败: ' + e.message, true);
        stopCurrentLoop(); // 执行失败时停止循环
    }
}

// 实际执行修炼的函数
async function executeCultivation() {
    try {
        if (typeof cultivation !== 'function') {
            showToast('修炼接口未加载', true);
            stopCurrentLoop();
            return;
        }

        const result = cultivation();
        showToast(result);
        refreshUserInfo();
        addToLog(`修炼结果: ${result}`);
    } catch (e) {
        showToast('修炼失败: ' + e.message, true);
        stopCurrentLoop();
    }
}

// 实际执行做工的函数
async function executeGetGold() {
    try {
        const result = getGold();
        showToast(result);
        refreshUserInfo();
        addToLog(`做工结果: ${result}`);
    } catch (e) {
        showToast('做工失败: ' + e.message, true);
        stopCurrentLoop();
    }
}
