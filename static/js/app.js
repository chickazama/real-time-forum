import LoginForm from "./components/LoginForm.js";
import Post from "./components/Post.js";
import SignupForm from "./components/SignupForm.js";
import { renderLoginForm, renderSignupForm } from "./render.js";

customElements.define("signup-form", SignupForm);
customElements.define("login-form", LoginForm);
customElements.define("post-element", Post);

let user;
const posts = new Map();

window.addEventListener("load", async () => {
    let res = await fetch("/api/user");
    if (res.status !== 200) {
        renderLoginForm();
        return;
    }
    let data = await res.json();
    main.innerText = `Welcome ${data.nickname}`;
    let logout = document.createElement("a");
    logout.innerText = "Logout";
    logout.href="/logout";
    main.appendChild(logout);
    await getPostsAsync();

});

async function getPostsAsync() {
    const res = await fetch("/api/posts");
    if (res.status !== 200) {
        console.log(res.statusText);
        return;
    }
    const data = await res.json();
    console.log(data);
    for (const item of data) {
        console.log(item);
        const post = new Post(item, null);
        main.appendChild(post);
        posts.set(item.id, post);
    }
}