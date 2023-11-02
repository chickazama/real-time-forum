import { renderSignupForm } from "../render.js";

export default class LoginForm extends HTMLElement {
    // Fields
    static observedAttributes = ["active"];
    shadowRoot;

    // Methods
    constructor() {
        super();
        let content = `
        <link type="text/css" rel="stylesheet" href="./static/css/forms.css" />
        <form class="form" action="/login" method="post">
            <span class="title">Login</span>
            <div class="input-field">
                <label for="input-nickname">Nickname</label>
                <input type="text" name="nickname" id="input-nickname" required="true" />
            </div>
            <div class="input-field">
            <label for="input-password">Password</label>
                <input type="password" name="password" id="input-password" required="true" />
            </div>
            <button type="submit">Log In</button>
        </form>
        <div>
            <button id="redirect">Go to Sign Up</button>
        </div>`;
        this.template = document.createElement("template");
        this.template.innerHTML = content;
        this.shadowRoot = this.attachShadow({ mode: "open" });
        this.shadowRoot.appendChild(this.template.content.cloneNode(true));
    }

    connectedCallback() {
        console.log("LoginForm added to page.");
        // this.shadowRoot.addEventListener("click", loginHandler);
        this.shadowRoot.addEventListener("click", redirectHandler);
    }

    disconnectedCallback() {
        console.log("LoginForm removed from page.");
    }

    attributeChangedCallback(name, oldValue, newValue) {
        console.log(`Attribute ${name} has changed.`);
        console.log(`${oldValue} -> ${newValue}`);
    }
}

function redirectHandler(event) {
    if (event.target.id != "redirect") {
        return;
    }
    const root = event.target.getRootNode();
    const host = root.host;
    host.setAttribute("active", "false");
    console.log("Login Redirect");
    renderSignupForm();
}