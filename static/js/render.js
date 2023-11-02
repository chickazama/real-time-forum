import LoginForm from "./components/LoginForm.js";
import SignupForm from "./components/SignupForm.js";

const main = document.getElementById("main");

export function renderSignupForm() {
    main.innerHTML = "";
    main.appendChild(new SignupForm());
}

export function renderLoginForm() {
    main.innerHTML = "";
    main.appendChild(new LoginForm());
}