import { renderChat } from "../app.js";

export default class Contact extends HTMLElement {
    shadowRoot;
    User;
    Data;

    constructor(user, data) {
        super();
        this.User = user;
        this.Data = data;
        const content =
        `
        <link type="text/css" rel="stylesheet" href="./static/css/chat.css" />
        <div class="contact" id="contact">
            <h5>${data.nickname}</h5>
        </div>
        `;
        const template = document.createElement("template");
        template.innerHTML = content;
        this.shadowRoot = this.attachShadow({mode: "open"});
        this.shadowRoot.appendChild(template.content.cloneNode(true));
    }

    connectedCallback() {
        console.log("Contact added to page.");
        this.shadowRoot.addEventListener("click", selectChatHandler);
    }

    disconnectedCallback() {
        console.log("Contact removed from page.");
    }
}

function selectChatHandler(event) {
    const root = event.target.getRootNode();
    const host = root.host;
    console.log(host.Data);
    renderChat(host.User, host.Data);
}