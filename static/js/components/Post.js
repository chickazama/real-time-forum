import { buildTimeString } from "../app.js";

const codeNewComment = 5;

export default class Post extends HTMLElement {
    static observedAttributes = ["active"];
    shadowRoot;
    User;
    Data;
    Comments;
    Socket;
    constructor(user, data, socket) {
        super();
        this.User = user;
        this.Data = data;
        this.Socket = socket;
        const timestamp = buildTimeString(data.timestamp);
        const content =
        `
        <link type="text/css" rel="stylesheet" href="./static/css/posts.css" />
        <div class="post" id="post">
            <h1>${data.content}</h1>
            <p>${data.categories}</p>
            <p>${data.author}</p>
            <p>${timestamp}</p>
            <div id="send-comment-container">
                <input type="text" id="comment-content" placeholder="Type a comment..." />
                <button id="send-comment" type="button">Send Comment</button>
            </div>
            <button id="display-comments-button">Show Comments</button>
            <div class="comments-container" id="comments-container"></div>
        </div>
        `;
        const template = document.createElement("template");
        template.innerHTML = content;
        this.shadowRoot = this.attachShadow({mode: "open"});
        this.shadowRoot.appendChild(template.content.cloneNode(true));
    }

    connectedCallback() {
        console.log("Post component added to DOM");
        this.shadowRoot.addEventListener("click", showCommentsHandler);
        this.shadowRoot.addEventListener("click", sendCommentHandler);
    }

    attributeChangedCallback(name, oldValue, newValue) {
        // console.log(`${name}: ${oldValue} -> ${newValue}`);
        switch (name) {
            case "active":
                switch(newValue) {
                    case "true":
                        if (oldValue != "refresh") {
                            this.shadowRoot.removeEventListener("click", showCommentsHandler);
                            this.showCommentsAsync().then( () => {
                                this.shadowRoot.addEventListener("click", hideCommentsHandler);
                            });
                        }
                        break;
                    case "false":
                        this.hideComments();
                        this.shadowRoot.removeEventListener("click", hideCommentsHandler);
                        this.shadowRoot.addEventListener("click", showCommentsHandler);
                        break;
                    case "refresh":
                        this.hideComments();
                        this.showCommentsAsync();
                        this.setAttribute("active", "true");
                        break;
                }
                break;
        }
    }

    async showCommentsAsync() {
        const commentsButton = this.shadowRoot.getElementById("display-comments-button");
        commentsButton.innerText = "Hide Comments";
        const commentsContainer = this.shadowRoot.getElementById("comments-container");
        if (!this.Comments) {
            this.Comments = [];
            const data = await this.getCommentsAsync(this.Data.id);
            if (!data) {
                return;
            }
            for (const item of data) {
                this.Comments.push(item);
            }
        } else {
            console.log("Comments already retrieved");
        }
        for (const comment of this.Comments) {
            const div = document.createElement("div");
            div.classList.add("comment");
            div.innerHTML =
            `
            <h3>${comment.author}</h3>
            <p>${comment.content}</h1>
            <p>${buildTimeString(comment.timestamp)}</p>
            `;
            commentsContainer.appendChild(div);
        }
    }

    hideComments() {
        const commentsButton = this.shadowRoot.getElementById("display-comments-button");
        commentsButton.innerText = "Show Comments";
        const commentsContainer = this.shadowRoot.getElementById("comments-container");
        commentsContainer.innerHTML = "";
    }

    async addCommentAsync(comment) {
        if (!this.Comments) {
            this.Comments = [];
            const data = await this.getCommentsAsync(this.Data.id);
            if (!data) {
                return;
            }
            for (const item of data) {
                this.Comments.push(item);
            }
            return
        }
        this.Comments.push(comment);
    }

    async getCommentsAsync(postID) {
        try {
            const response = await fetch("/api/comments", {
                method: "POST",
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                    "Accept": "application/json"
                },
                body: new URLSearchParams( 
                    {
                        "postID":  postID
                    }
                )
            });
            const result = await response.json();
            console.log("Success:", result);
            return result;
        } catch (error) {
            console.error("Error:", error);
            return null;
        }
    }
}

function showCommentsHandler(event) {
    if (event.target.id != "display-comments-button") {
        return;
    }
    console.log(`'showCommentsHandler' fired: ${event}`);
    const root = event.target.getRootNode();
    const host = root.host;
    host.setAttribute("active", "true");
}

function hideCommentsHandler(event) {
    if (event.target.id != "display-comments-button") {
        return;
    }
    console.log(`'showCommentsHandler' fired: ${event}`);
    const root = event.target.getRootNode();
    const host = root.host;
    host.setAttribute("active", "false");
}

function sendCommentHandler(event) {
    if (event.target.id != "send-comment") {
        return;
    }
    const root = event.target.getRootNode();
    const host = root.host;
    const input = root.getElementById("comment-content");
    const contentData = input.value;
    if (contentData.length <= 0) {
        alert("Comment must have content");
        return;
    }
    let dummyData = {
        postID: host.Data.id,
        authorID: host.User.id,
        author: host.User.nickname,
        content: contentData,
        timestamp: Math.floor(Date.now() / 1000)
    };

    let body = {
        code: codeNewComment,
        data: dummyData
    };
    host.Socket.send(JSON.stringify(body));
    input.value = "";
}