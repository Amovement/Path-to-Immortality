
// 挑战相关函数
async function listChallengeBtn() {
    try {
        // 1. 先检查Wasm接口是否加载成功
        if (typeof listChallenge !== 'function') {
            showToast('挑战接口未加载，请稍后重试', true);
            return;
        }

        // 2. 调用Wasm接口获取数据
        const rawData = listChallenge();
        let challengeData = null;

        // 3. 解析JSON数据（处理可能的解析错误）
        try {
            challengeData = JSON.parse(rawData);
            console.log('获取到的挑战数据:', challengeData); // 调试用，可删除
        } catch (jsonErr) {
            showToast('挑战数据解析失败', true);
            console.error('JSON解析错误:', jsonErr, '原始数据:', rawData);
            return;
        }

        // 4. 提取核心的list数组（关键修复：适配{list: Array}格式）
        let challengeList = [];
        // 先判断是否是对象，且包含list属性，且list是数组
        if (typeof challengeData === 'object' && challengeData !== null && Array.isArray(challengeData.list)) {
            challengeList = challengeData.list; // 从对象中提取真正的数组
        } else {
            showToast('挑战数据格式异常', true);
            console.warn('预期数据格式是{list: Array}，实际是:', challengeData);
            return;
        }

        // 5. 渲染挑战列表
        const challengeListElement = document.getElementById('challengeList');
        if (challengeList.length === 0) {
            challengeListElement.innerHTML = `
                <div class="text-center text-dark-400 py-8">
                    <i class="fa fa-frown-o text-xl mb-2"></i>
                    <p>暂无可用挑战</p>
                </div>
            `;
            return;
        }

        // 6. 循环生成挑战项（给每个属性加默认值，避免缺失属性报错）
        let html = '';
        challengeList.forEach((challenge, index) => {
            const id = challenge.id || `${index}`; // 用索引当默认ID，避免空值
            const name = challenge.title || `未知挑战-${index + 1}`;
            const reward = challenge.reward || '无奖励';
            const description = challenge.description || '暂无挑战描述';

            html += `
                <div class="border border-dark-300 rounded-lg p-4 hover:border-primary/50 transition-colors">
                    <div class="flex justify-between items-start">
                        <div>
                            <h3>${name}</h3>
                            <p class="text-dark-400 text-sm">奖励: ${reward}</p>
                        </div>
                        <button onclick="joinChallengeBtn(${id})"
                            class=" text-accent px-3 py-1 rounded hover:bg-accent/30 transition-colors text-sm">
                            接受挑战
                        </button>
                    </div>
                    <p class="mt-2 text-dark-400 text-sm whitespace-pre-line">${description}</p>
                </div>
            `;
        });

        challengeListElement.innerHTML = html;
        showToast(`已加载 ${challengeList.length} 个挑战`);
    } catch (e) {
        showToast('获取挑战列表失败: ' + e.message, true);
        console.error('挑战列表加载异常:', e);
    }
}

async function joinChallengeBtn(challengeId) {
    try {
        if (typeof joinChallenge !== 'function') {
            showToast('接口未加载', true);
            return;
        }

        const result = joinChallenge(challengeId);
        console.log('加入挑战结果:', result);
        refreshUserInfo();

        if (result.log) {
            addToLog(result.log);
            addToLog(result.msg)
        }
        showToast(result.msg);
    } catch (e) {
        showToast('加入挑战失败: ' + e.message, true);
    }
}

// 清空日志函数
function clearChallengeLog() {
    const logElement = document.getElementById('challengeLog');
    // 保留"暂无记录"的提示文本
    logElement.innerHTML = '<p class="italic">暂无记录</p>';
    showToast('日志已清空');
    addToLog('日志已清空');
}