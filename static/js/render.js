import LoginForm from "./components/LoginForm.js";
import SignupForm from "./components/SignupForm.js";

const header = document.getElementById("header");
const main = document.getElementById("main");

export function renderSignupForm() {
    header.innerHTML = "";
    main.innerHTML = "";
    main.appendChild(new SignupForm());
}

export function renderLoginForm() {
    header.innerHTML = "";
    main.innerHTML = "";
    main.appendChild(new LoginForm());
}