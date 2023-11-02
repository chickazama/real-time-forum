import LoginForm from "./components/LoginForm.js";
import SignupForm from "./components/SignupForm.js";
import { renderLoginForm, renderSignupForm } from "./render.js";

customElements.define("signup-form", SignupForm);
customElements.define("login-form", LoginForm);


window.addEventListener("load", async () => {
    let res = await fetch("/api/user");
    if (res.status == 200) {
        let json = await res.json();
        main.innerText = `Welcome ${json.nickname}`;
        let logout = document.createElement("a");
        logout.innerText = "Logout";
        logout.href="/logout";
        main.appendChild(logout);
    } else {
        renderLoginForm();
    }
})