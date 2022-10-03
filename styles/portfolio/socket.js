class MySocket {
    constructor() {
        this.mysocket = null;

        // this.messageHistory = document.querySelector('.messageHistory');

    }
    showMessage(text){
        console.log(text)
        // const jtext = JSON.parse(text)
        // if (jtext != null){
        //     jtext.forEach((element) => {
        //         this.addMessage(element['MessageText'],element['SenderId'])
        //     })
        // }
    }
    send(type, data, cleanedEl){
        console.log({
            req_type: type,
            user_balance_coins :data
        })
        if (type == "buy"){
            this.mysocket.send(JSON.stringify({
                req_type: type,
                user_balance_coins :data
            }))
            cleanedEl.value = ""
        }
    }
    connectSocket(){
        const socket = new WebSocket("ws://localhost:8080/api/ws/" + this.getCookie('public_address_token'));
        this.mysocket = socket;

        socket.onmessage = (e)=>{
            this.showMessage(e.data);
        }
        socket.onopen =  ()=> {
        };
        socket.onclose = ()=>{
            this.reconnect()
        }
    }
    getCookie(cookieName) {
        let cookie = {};
        document.cookie.split(';').forEach(function(el) {
            let [key,value] = el.split('=');
            cookie[key.trim()] = value;
        })
        return cookie[cookieName];
    }
    // addMessage(text = '', author = 0) {
    //     const newMessage = document.createElement('div')
    //     newMessage.classList.add('message')
    //     if (author != 1) {
    //         newMessage.classList.add('green')
    //     }
    //     newMessage.innerText = text
    //     const messageContainer = document.createElement('div')
    //     messageContainer.classList.add('messageContainer')
    //     messageContainer.appendChild(newMessage)
    //     this.messageHistory.appendChild(messageContainer)
    // }
    async reconnect() {
        await delay(5000);
        this.connectSocket()
    }
}
function delay(time) {
    return new Promise(resolve => setTimeout(resolve, time));
}
