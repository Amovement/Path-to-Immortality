// 加载商品
function openShopModal() {
    // 获取商品列表
    loadAllGoods(); // 加载所有商品
}

// 从Wasm加载所有商品
function loadAllGoods() {
    try {
        // 调用Wasm暴露的获取商品列表的函数
        if (typeof getGoodsList !== 'function') {
            showToast('商品接口未加载', true);
            return;
        }

        const goodsListStr = getGoodsList();
        const goodsList = JSON.parse(goodsListStr);

        // 检查是否是数组
        if (!Array.isArray(goodsList)) {
            showToast('商品数据格式错误', true);
            console.error('商品数据不是数组:', goodsList);
            return;
        }

        renderGoodsList(goodsList);
    } catch (e) {
        showToast('加载商品失败: ' + e.message, true);
        console.error('加载商品错误:', e);
        // 显示错误状态
        document.getElementById('challengeList').innerHTML = `
                    <div class="text-center text-dark-400 py-10">
                        <i class="fa fa-exclamation-triangle text-danger text-xl mb-2"></i>
                        <p>无法加载商品，请重试</p>
                        <button onclick="loadAllGoods()" class="mt-2 text-primary text-sm hover:underline">
                            重试
                        </button>
                    </div>
                `;
    }
}

// 渲染商品列表
function renderGoodsList(goodsList) {
    const goodsListElement = document.getElementById('challengeList');

    if (goodsList.length === 0) {
        goodsListElement.innerHTML = `
                    <div class="text-center text-dark-400 py-10">
                        <i class="fa fa-frown-o text-xl mb-2"></i>
                        <p>当前暂无商品</p>
                    </div>
                `;
        return;
    }

    let html = '';
    goodsList.forEach(goods => {
        // 从Goods结构体获取属性，添加默认值
        const id = goods.uuid || 0;
        const name = goods.name || '未知商品';
        const description = goods.description || '暂无描述';
        const price = goods.price || 0;

        // 根据商品名称自动选择图标（可根据实际情况调整）
        let iconClass = 'fa-cube'; // 默认图标

        html += `
            <div class="flex flex-col md:flex-row gap-0.2 p-0.2 bg-dark-100 rounded-lg border border-dark-300 shop-item-hover">
                <div class="w-full md:w-1 h-1 bg-dark-300 rounded flex items-center justify-center relative">
                </div>
                <div class="flex-1">
                    <!-- 使用 justify-between 确保标题左对齐，金币标签右对齐 -->
                    <div class="flex justify-between items-center">
                        <h3>${name}</h3>
                        <!-- 添加 whitespace-nowrap 防止价格标签换行 -->
                        <span class="text-primary text-sm px-1 py-0.5 rounded bg-primary/10 whitespace-nowrap">
                            ${price} 金
                        </span>
                    </div>
                    <p class="text-dark-400 md:text-sm mt-0.5 mb-0.5">
                        ${description}
                    </p>
                    <div class="mt-1 flex justify-end">
                        <button onclick="buyGoodsBtn(${id})" class=" text-gray-100 px-0.5 py-0.5 rounded hover:bg-secondary text-sm transition-colors">
                            <i class="fa fa-shopping-cart mr-1"></i>购买
                        </button>
                    </div>
                </div>
            </div>`;
    });

    goodsListElement.innerHTML = html;
    showToast(`已加载 ${goodsList.length} 个商品`);
}

// 购买商品
function buyGoodsBtn(goodsId) {
    try {
        // 调用Wasm的购买接口
        if (typeof buyGoods !== 'function') {
            console.log('buyGoods is not a function')
            return;
        }

        // 调用购买函数，传入商品ID
        const result = buyGoods(goodsId);

        // 返回格式为 '"金币不足"'
        addToLog(`购买结果: ${result}`);
        refreshUserInfo(); // 刷新用户信息
        showToast(result,true);
    } catch (e) {
        showToast('购买失败: ' + e.message, true);
        console.error('购买商品错误:', e);
    }
}