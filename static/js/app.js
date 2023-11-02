import SignupForm from "./components/SignupForm.js";

customElements.define("signup-form", SignupForm);

const main = document.getElementById("main");

window.addEventListener("load", () => {
    main.appendChild(new SignupForm());
})