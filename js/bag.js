function openBag(){
    try {
        // 调用Wasm暴露的获取商品列表的函数
        if (typeof getBag !== 'function') {
            return;
        }

        const backpackJson = getBag()
        const backpackData = JSON.parse(backpackJson);
        renderBackpackItems(backpackData);


    } catch (e) {
        showToast('加载背包失败: ' + e.message, true);
        console.error('加载背包失败:', e);
        // 显示错误状态
        document.getElementById('challengeList').innerHTML = `
                    <div class="text-center text-dark-400 py-10">
                        <i class="fa fa-exclamation-triangle text-danger text-xl mb-2"></i>
                        <p>无法加载背包，请重试</p>
                        <button onclick="openBag()" class="mt-2 text-primary text-sm hover:underline">
                            重试
                        </button>
                    </div>
                `;
    }
}

// 渲染背包物品列表
function renderBackpackItems(backpackData) {
    const backpackElement = document.getElementById('challengeList'); // 假设背包容器ID为backpackList

    const itemsList = backpackData.Items;

    // 空背包状态
    if (itemsList==null || itemsList.length === 0) {
        backpackElement.innerHTML = `
            <div class="text-center text-dark-400 py-10">
                <i class="fa fa-briefcase text-xl mb-2"></i>
                <p>背包空空如也</p>
            </div>
        `;
        return;
    }

    let itemCnt = 0;
    for (let i = 0; i < itemsList.length; i++) {
        const item = itemsList[i];
        const type = item.type || 0;
        if (type === 1) {
            continue;
        }
        itemCnt++;
    }

    // 空背包状态
    if (itemCnt === 0) {
        backpackElement.innerHTML = `
            <div class="text-center text-dark-400 py-10">
                <i class="fa fa-briefcase text-xl mb-2"></i>
                <p>背包空空如也</p>
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
        let equipDescription = "";
        if (type ===1){ // 法器不显示在背包里
            return
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
                        ${type === 0 ? '<span class="text-primary text-sm px-1 py-0.5 rounded bg-primary/10">消耗品</span>' : ''}
                        ${type === 1 ? '<span class="text-primary text-sm px-1 py-0.5 rounded bg-primary/20">法器</span>' : ''}
                        ${type === 2 ? '<span class="text-primary text-sm px-1 py-0.5 rounded bg-primary/30">材料</span>' : ''}
                    </div>
                    <!-- 减少描述文本的上下间距 -->
                    <p class="text-dark-400 md:text-sm mt-0.5 mb-0.5 ">
                        ${description}
                        ${equipDescription===''?'' : '</br>'+equipDescription}
                    </p>
                    <div class="mt-1 flex justify-end">
                        <button onclick="useBagItemBtn(${id})" class=" text-gray-100 px-0.5 py-0.5 rounded hover:bg-secondary text-sm transition-colors">
                            <i class="fa fa-plus-circle"></i>使用
                        </button>
                    </div>
                </div>
            </div>
        `;
    });

    backpackElement.innerHTML = html;
    showToast(`背包中有 ${itemCnt} 个物品`);
}

function useBagItemBtn(item) {
    const result = useBagItem(item);
    showToast(result);
    addToLog(result);
    openBag();
    refreshUserInfo();
}