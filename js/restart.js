// 显示确认模态框
function showRestartConfirm() {
    const modal = document.getElementById('restartConfirmModal');
    modal.classList.remove('opacity-0', 'pointer-events-none');
}

// 隐藏确认模态框
function hideRestartConfirm() {
    const modal = document.getElementById('restartConfirmModal');
    modal.classList.add('opacity-0', 'pointer-events-none');
}

// 确认转生（原restartBtn逻辑）
async function confirmRestart() {
    // 先隐藏确认框
    hideRestartConfirm();

    // 执行转生逻辑
    const result = restart();
    refreshUserInfo();
    showToast(result);
    addToLog(`转生: ${result}`);

    // 隐藏转生按钮
    const btn = document.getElementById('restartB');
    if (btn) {
        btn.classList.add('hidden');
    }
}