const isMetaMaskInstalled = () => {
    //Have to check the ethereum binding on the window object to see if it's installed
    const { ethereum } = window;
    return Boolean(ethereum && ethereum.isMetaMask);
};
async function onClickConnect() {
    try {
        const web3 = new Web3(window.ethereum);
        //получаем адрес
        const accounts = await ethereum.request({ method: 'eth_requestAccounts' });
        document.querySelector(".account_data").innerHTML= accounts[0]
        console.log("адрес кошелька: " + accounts[0])
        const message = 'Very Message Such Wow';
        console.log("Сообщение на подписание: " + message)
        const from = accounts[0];
        //подписываем текст
        const sign = await ethereum.request({
            method: 'personal_sign',
            params: [message, from, 'Random text'],
        });
        console.log("Подписанное сообщение: " + sign)
        //проверяем подлинность подписи, если всё верно recoveredAddr == accounts[0]
        const recoveredAddr = web3.eth.accounts.recover(message, sign);
        console.log('Раскодированный адрес на основе базового сообщения и полученного: ' + recoveredAddr);
    } catch (error) {
        console.error(error);
    }
};