import { buildTimeString, renderUsers } from "../app.js";

const codeNewDirectMessage = 6;


export default class Chat extends HTMLElement {
    static observedAttributes = ["active"];
    shadowRoot;
    Sender;
    Target;
    Messages;
    Socket;
    Offset = 10;   

    constructor(sender, target, socket) {
        super();
        this.Sender = sender;
        this.Target = target;
        this.Socket = socket;
        const content =
        `
        <link type="text/css" rel="stylesheet" href="./static/css/chat.css" />
        <button id="redirect" type="button">Back to Contacts</button>
        <h5>${target.nickname}</h5>
        <div id="messages-container"></div>
        <div id="send-message-container">
            <input type="text" id="message-content" placeholder="Type a message..." />
            <button id="send-message" type="button">Send Message</button>
        </div>
        `;
        const template = document.createElement("template");
        template.innerHTML = content;
        this.shadowRoot = this.attachShadow({mode: "open"});
        this.shadowRoot.appendChild(template.content.cloneNode(true));
    }

    connectedCallback() {
        console.log("Chat component added to DOM");
        this.shadowRoot.addEventListener("click", sendMessageHandler);
        this.shadowRoot.addEventListener("click", redirectHandler);
        const container = this.shadowRoot.getElementById("messages-container");
        container.addEventListener("scroll", () => {
            this.debounce(this.scrollReloadHandler, 100, container); 
        });
        if (!this.Messages) {
            this.getMessagesAsync(this.Target.id, 0).then( (res) => {
                console.log(this.Offset);
                this.Messages = [];
                for (const item of res) {
                    this.Messages.push(item);
                }
                this.displayMessages();
            })
        } else {
            this.displayMessages();
        }
        
    }

    disconnectedCallback() {
        console.log("Chat component removed from DOM");
        this.shadowRoot.removeEventListener("click", sendMessageHandler);
        this.shadowRoot.addEventListener("click", redirectHandler);
    }

    displayMessages() {
        const messagesContainer = this.shadowRoot.getElementById("messages-container");
        messagesContainer.innerHTML = "";
        for (const message of this.Messages) {
            const div = document.createElement("div");
            div.classList.add("message");
            div.innerHTML =
            `
            <h5>${message.author}</h5>
            <h3>${message.content}</h3>
            <p>${buildTimeString(message.timestamp)}</p>
            `;
            messagesContainer.appendChild(div);
        }
    }

    async addMessageAsync(message) {
        if (!this.Messages) {
            this.Messages = [];
            const data = await this.getMessagesAsync(this.Target.id);
            if (!data) {
                return;
            }
            for (const item of data) {
                this.Messages.push(item);
            }
            return
        }
        const msg = [message]
        this.Messages = msg.concat(this.Messages);
    }

    scrollReloadHandler(container) {
        const root = container.getRootNode();
        const host = root.host;
        if(Math.abs((container.offsetHeight - container.scrollHeight)- container.scrollTop) < 5) {
            console.log("Scrolled to top");
            host.getMessagesAsync(host.Target.id, host.Offset).then((data) => {
                for (const message of data) {
                    host.Messages.push(message);
                    const div = document.createElement("div");
                    div.classList.add("message");
                    div.innerHTML =
                    `
                    <h5>${message.author}</h5>
                    <h3>${message.content}</h3>
                    <p>${buildTimeString(message.timestamp)}</p>
                    `;
                    container.appendChild(div);
                }
            })
            host.Offset += 10;
        }
    }

    async getMessagesAsync(targetID, offset) {
        try {
            const response = await fetch("/api/messages?" + new URLSearchParams({
                limit: 10,
                offset: offset,
            }), {
                method: "POST",
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                    "Accept": "application/json"
                },
                body: new URLSearchParams( 
                    {
                        "targetID":  targetID
                    }
                )
            });
            const result = await response.json();
            console.log("Success:", result);
            return result;
        } catch (error) {
            console.error("Error:", error);
            return null;
        }
    }

    debounce(method, delay, args) {
        clearTimeout(method._tId);
        method._tId= setTimeout(function(){
            method(args);
        }, delay);
    }
}

function redirectHandler(event) {
    if (event.target.id != "redirect") {
        return;
    }
    renderUsers();
}

function sendMessageHandler(event) {
    if (event.target.id != "send-message") {
        return;
    }
    const root = event.target.getRootNode();
    const host = root.host;
    const input = root.getElementById("message-content");
    const contentData = input.value;
    if (contentData.length <= 0) {
        alert("Message must have content");
        return;
    }
    let dummyData = {
        senderID: host.Sender.id,
        targetID: host.Target.id,
        author: host.Sender.nickname,
        content: contentData,
        timestamp: Math.floor(Date.now() / 1000)
    };

    let body = {
        code: codeNewDirectMessage,
        data: dummyData
    };
    host.Socket.send(JSON.stringify(body));
}