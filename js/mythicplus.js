function listMythicPlusBtn(){
    try {
        // 1. 先检查Wasm接口是否加载成功
        if (typeof getMythicInfo !== 'function') {
            showToast('秘境接口未加载，请稍后重试', true);
            return;
        }

        // 2. 调用Wasm接口获取数据
        const rawData = getMythicInfo();
        let mythicPlusData = null;

        // 3. 解析JSON数据（处理可能的解析错误）
        try {
            mythicPlusData = JSON.parse(rawData);
            console.log('获取到的秘境数据:', mythicPlusData); // 调试用，可删除
        } catch (jsonErr) {
            showToast('数据解析失败', true);
            console.error('JSON解析错误:', jsonErr, '原始数据:', rawData);
            return;
        }

        // 渲染列表
        const challengeListElement = document.getElementById('challengeList');
        if (mythicPlusData.monsters.length === 0) {
            challengeListElement.innerHTML = `
                <div class="text-center text-dark-400 py-8">
                    <i class="fa fa-frown-o text-xl mb-2"></i>
                    <p>秘境数据异常</p>
                </div>
            `;
            return;
        }
        const monsterData = mythicPlusData.monsters;
        //
        // 循环渲染怪物（给每个属性加默认值，避免缺失属性报错）
        let html = `
            <p class="text-accent px-3 py-1 rounded transition-colors text-sm">当前秘境: ${mythicPlusData.level} 层</p>
            <button onclick='joinMythicBtn()' class="text-accent px-3 py-1 rounded hover:bg-accent/30 transition-colors text-sm">挑战当前层</button>
            <button onclick='lowerTheMythicPlusBtn()' class="text-accent px-3 py-1 rounded hover:bg-accent/30 transition-colors text-sm">秘境降级</button>
            
            <!-- 词缀解释悬浮容器 -->
            <div class="affix-tooltip-container relative inline-block">
                <div class="text-accent cursor-help px-3 py-1 rounded transition-colors text-sm">词缀解释</div>
                <div class="affix-tooltip absolute z-10 invisible opacity-0 transition-all duration-300 bg-dark-800 text-light p-3 rounded-lg shadow-lg w-64 -mt-2 left-0">
                    <ul class="text-sm mt-2 space-y-1">
                        ${mythicPlusData.description}
                    </ul>
                </div>
            </div>
        `;
        monsterData.forEach((monster, index) => {
            const attack = monster.attack || 0;
            const cultivation = monster.cultivation || '灵智未启';
            const defense = monster.defense || 0;
            const hp = monster.hp || 0;
            const speed = monster.speed || 0;
            const hpLimit = monster.hpLimit || 0;
            const name = monster.name || `未知怪物-${index + 1}`;
            const special = monster.special || [];
            const level = monster.level || 0;

            html += `
                <div class="border border-dark-300 rounded-lg p-4 hover:border-primary/50 transition-colors">
                    <div class="flex justify-between items-start">
                        <div>
                            <h3>${name}</h3>
                            <p class="mt-2 text-dark-400 text-sm whitespace-pre-line">${cultivation}</p>
                        </div>
                    </div>
                    <p class="text-dark-400 text-sm">生命: ${hp}/${hpLimit}    攻击: ${attack}    防御: ${defense}    速度: ${speed} </p>
                    <p class="text-dark-400 text-sm">特殊词缀: ${special}</p>
                </div>
            `;
        });

        challengeListElement.innerHTML = html;
        showToast(`已加载秘境列表`);
    } catch (e) {
        showToast('获取秘境列表失败: ' + e.message, true);
        console.error('秘境列表加载异常:', e);
    }
}

function lowerTheMythicPlusBtn(){
    try{
        const rawData = lowerTheMythicPlus();
        addToLog(rawData);
        listMythicPlusBtn();
    } catch (e) {
        showToast('获取秘境列表失败: ' + e.message, true);
        console.error('秘境列表加载异常:', e);
    }
}

function joinMythicBtn(){
    try{
        const result = joinMythic();
        refreshUserInfo();

        if (result.log) {
            addToLog(result.log);
            addToLog(result.msg)
        }
        showToast(result.msg);
        listMythicPlusBtn();
    } catch (e) {
        showToast('失败: ' + e.message, true);
        console.error('加载异常:', e);
    }
}