
function listEquip(){
    try {
        // 调用Wasm暴露的获取背包列表的函数
        if (typeof getBag !== 'function') {
            return;
        }
        const backpackJson = getBag()
        const backpackData = JSON.parse(backpackJson);
        renderEquipItems(backpackData);
    } catch (e) {
        showToast('加载法器失败: ' + e.message, true);
        console.error('加载法器失败:', e);
        // 显示错误状态
        document.getElementById('challengeList').innerHTML = `
                    <div class="text-center text-dark-400 py-10">
                        <i class="fa fa-exclamation-triangle text-danger text-xl mb-2"></i>
                        <p>无法加载法器，请重试</p>
                        <button onclick="listEquip()" class="mt-2 text-primary text-sm hover:underline">
                            重试
                        </button>
                    </div>
                `;
    }
}

// 渲染背包物品列表
function renderEquipItems(backpackData) {
    const backpackElement = document.getElementById('challengeList');

    const itemsList = backpackData.Items;

    // 空背包状态
    if (itemsList==null || itemsList.length === 0) {
        backpackElement.innerHTML = `
            <div class="text-center text-dark-400 py-10">
                <i class="fa fa-briefcase text-xl mb-2"></i>
                <p>法器空空如也</p>
            </div>
        `;
        return;
    }

    let itemCnt = 0;
    for (let i = 0; i < itemsList.length; i++) {
        const item = itemsList[i];
        const type = item.type || 0;
        if (type === 1) {
            itemCnt++;
        }
    }

    if (itemCnt === 0) {
        backpackElement.innerHTML = `
            <div class="text-center text-dark-400 py-10">
                <i class="fa fa-briefcase text-xl mb-2"></i>
                <p>法器空空如也</p>
            </div>
        `;
        return;
    }

    // 渲染物品列表
    let html = '';
    itemsList.forEach(item => {
        // 从物品数据获取属性，添加默认值
        const id = item.uuid || 0;
        const name = item.name || '未知物品';
        const description = item.description || '暂无描述';
        const count = item.count || 0;
        const type = item.type || 0;
        const equipInfo = item.equipInfo || null;
        const status = item.status || null;
        let equipDescription = "";
        if (type ===1){ // 法器
            equipDescription = "法器属性 攻击: " + equipInfo.Attack + " 防御: " + equipInfo.Defense + " 速度: " + equipInfo.Speed + " 体魄: " + equipInfo.Hp;
        }else{
            return  ;
        }

        // 物品数量标签（超过1个才显示）
        const countBadge = count > 1 ?
            `<span class="absolute -top-1 -right-1 bg-accent text-dark-100 text-xs rounded-full w-5 h-5 flex items-center justify-center">
                ${count}
            </span>` : '';

        html += `
            <div class="flex flex-col md:flex-row gap-0.2 p-0.2 bg-dark-100 rounded-lg border border-dark-300 shop-item-hover">
                <div class="w-full md:w-1 h-1 bg-dark-300 rounded flex items-center justify-center relative">
                    ${countBadge}
                </div>
                <div class="flex-1">
                    <div class="flex justify-between">
                        <h3 >${name}</h3>
                        ${equipInfo.Type === 0 ? '<span class="text-primary text-sm px-1 py-0.5 rounded bg-primary/10">头甲</span>' : ''}
                        ${equipInfo.Type === 1 ? '<span class="text-primary text-sm px-1 py-0.5 rounded bg-primary/10">胸甲</span>' : ''}
                        ${equipInfo.Type === 2 ? '<span class="text-primary text-sm px-1 py-0.5 rounded bg-primary/10">臂甲</span>' : ''}
                        ${equipInfo.Type === 3 ? '<span class="text-primary text-sm px-1 py-0.5 rounded bg-primary/10">腿甲</span>' : ''}
                        ${equipInfo.Type === 4 ? '<span class="text-primary text-sm px-1 py-0.5 rounded bg-primary/10">武器</span>' : ''}
                        ${equipInfo.Type === 5 ? '<span class="text-primary text-sm px-1 py-0.5 rounded bg-primary/10">饰品</span>' : ''}
                    </div>
                    <!-- 减少描述文本的上下间距 -->
                    <p class="text-dark-400 md:text-sm mt-0.5 mb-0.5 ">
                        ${description}
                        ${equipDescription===''?'' : '</br>'+equipDescription}
                    </p>
                    
                    ${status === 1 ? '<span class="text-primary text-sm px-1 py-0.5 rounded bg-primary/10">装备中</span>' : ''}
                    <div class="mt-1 flex justify-end">
                        <button onclick="useEquipItemBtn(${id})" class=" text-gray-100 px-0.5 py-0.5 rounded hover:bg-secondary text-sm transition-colors">
                            <i class="fa fa-plus-circle"></i> ${status === 1 ? '卸下' : '装备'}
                        </button>
                    </div>
                </div>
            </div>
        `;
    });

    backpackElement.innerHTML = html;
    showToast(`背包中有 ${itemCnt} 个法器`);
}

function useEquipItemBtn(item) {
    const result = useBagItem(item);
    showToast(result);
    addToLog(result);
    listEquip();
    refreshUserInfo();
}