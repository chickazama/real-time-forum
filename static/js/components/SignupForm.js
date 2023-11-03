import { renderLoginForm } from "../app.js";

export default class SignupForm extends HTMLElement {
    static observedAttributes = ["active"];
    shadowRoot;
    constructor() {
        super();
        const content = `
        <link type="text/css" rel="stylesheet" href="./static/css/forms.css" />
        <form class="form" action="/signup" method="post">
            <span class="title">Signup</span>
            <div class="input-field">
                <label for="input-nickname">Nickname</label>
                <input type="text" name="nickname" id="input-nickname" required="true" />
            </div>
            <div class="input-field">
                <label for="input-age">Age</label>
                <input type="number" name="age" id="input-age" required="true" />
            </div>
            <div class="input-field">
                <label for="input-gender">Gender</label>
                <select name="gender" id="input-gender" placeholder="Gender" required="true" />
                    <option value="" disabled selected>Select Gender</option>
                    <option value="Male">Male</option>
                    <option value="Female">Female</option>
                    <option value="Other">Other</option>
                    <option value="Prefer Not To Say">Prefer Not To Say</option>
                </select>
            </div>
            <div class="input-field">
                <label for="input-first-name">First Name</label>
                <input type="text" name="firstName" id="input-first-name" required="true" />
            </div>
            <div class="input-field">
                <label for="input-last-name">Last Name</label>
                <input type="text" name="lastName" id="input-last-name" required="true" />
            </div>
            <div class="input-field">
                <label for="input-email">Email</label>
                <input type="email" name="emailAddress" id="input-email" required="true" />
            </div>
            <div class="input-field">
                <label for="input-password">Password</label>
                <input type="password" name="password" id="input-password" required="true" />
            </div>
            <button type="submit">Sign Up</button>
        </form>
        <div>
            <button id="redirect">Go to Log In</button>
        </div>`;
        const template = document.createElement("template");
        template.innerHTML = content;
        this.shadowRoot = this.attachShadow({ mode: "open" });
        this.shadowRoot.appendChild(template.content.cloneNode(true));
    }

    connectedCallback() {
        console.log("Signup Form added to page.");
        this.shadowRoot.addEventListener("click", redirectHandler);
    }

    disconnectedCallback() {
        console.log("Signup Form removed from page.");
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
    renderLoginForm();
}