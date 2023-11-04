export default class Navbar extends HTMLElement {
    shadowRoot;
    User;
    constructor(user) {
        super();
        this.User = user;
        const content =
        `
        <link type="text/css" rel="stylesheet" href="./static/css/navbar.css" />
        <ul class="navbar">
            <li class="nav-item">
                <span class="nav-link">${user.nickname}</span>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/logout">Log Out</a>
            </li>
        </ul>
        `;
        const template = document.createElement("template");
        template.innerHTML = content;
        this.shadowRoot = this.attachShadow({mode: "open"});
        this.shadowRoot.appendChild(template.content.cloneNode(true));
    }

    connectedCallback() {
        console.log("Navbar element appended to page");
    }

    disconnectedCallback() {
        console.log("Navbar element removed from page");
    }
}

// function logoutHandler(event) {
//     if (event.target.id != "logout") {
//         return;
//     }
//     const root = event.target.getRootNode();
//     const host = root.host;
//     const res = 
// }