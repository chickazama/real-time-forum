import CreatePostForm from "./components/CreatePostForm.js";
import LoginForm from "./components/LoginForm.js";
import Post from "./components/Post.js";
import SignupForm from "./components/SignupForm.js";
import { renderLoginForm, renderSignupForm } from "./render.js";

customElements.define("signup-form", SignupForm);
customElements.define("create-post-form", CreatePostForm);
customElements.define("login-form", LoginForm);
customElements.define("post-element", Post);

const codeNewPost = 4;
const codeNewComment = 5;

const header = document.getElementById("header");
const main = document.getElementById("main");

let user;
let socket;
const posts = new Map();

window.addEventListener("load", async () => {
    let res = await fetch("/api/user");
    if (res.status !== 200) {
        renderLoginForm();
        return;
    }
    let data = await res.json();
    user = data;
    let logout = document.createElement("a");
    logout.innerText = "Logout";
    logout.href="/logout";
    header.appendChild(logout);
    setupWebSocket();
    let postForm = new CreatePostForm(user, socket);
    header.appendChild(postForm);
    let postsData = await getPostsAsync();
    for (const item of postsData) {
        console.log(item);
        const post = new Post(user, item, socket);
        main.appendChild(post);
        posts.set(item.id, post);
    }
});

async function getPostsAsync() {
    const res = await fetch("/api/posts");
    if (res.status !== 200) {
        console.log(res.statusText);
        return null;
    }
    const data = await res.json();
    return data;
}

function setupWebSocket() {
    socket = new WebSocket("ws://localhost:8080/websocket");
    socket.onopen = () => {
        console.log("Socket opened.");
    }
    socket.onclose = () => {
        console.log("Socket closed.");
    }
    socket.onmessage = async (e) => {
        // console.log(e.data);
        const message = JSON.parse(e.data);
        const body = JSON.parse(message.body);
        switch (body.code) {
            case codeNewPost:
                console.log("New Post - Data");
                console.log(body.data);
                const newPost = new Post(user, body.data, socket);
                posts.set(body.data.id, newPost);
                main.prepend(newPost);
                break;
            case codeNewComment:
                console.log("New Comment - Data:");
                const post = posts.get(body.data.postID);
                await post.addCommentAsync(body.data);
                if (post.getAttribute("active") == "true") {
                    post.setAttribute("active", "refresh");
                }
                break;
        }
    }
    socket.onerror = (e) => {
        console.log(e);
    }
}