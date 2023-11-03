const codeNewPost = 4;

export default class CreatePostForm extends HTMLElement {
    // Fields
    static observedAttributes = ["active"];
    shadowRoot;
    User;
    Socket;
    // Methods
    constructor(user, socket) {
        super();
        this.User = user;
        this.Socket = socket;
        let content = `
        <link type="text/css" rel="stylesheet" href="./static/css/forms.css" />
        <form class="form">
            <div class="categories">
                <input type="checkbox" id="category1" name="category1" value="Go" />
                <label for="category1"> Go</label>
                <input type="checkbox" id="category2" name="category2" value="JavaScript" />
                <label for="category2"> JavaScript</label>
                <input type="checkbox" id="category3" name="category3" value="Rust" />
                <label for="category3"> Rust</label>
            </div>
            <div class="content">
                <textarea class="text-area" id="text-area" placeholder="Write text here..."></textarea>
            </div>
            <div class="post-button-container">
                <button id="send-post" type="button">Post</button>
            </div>
        </form>`;
        const template = document.createElement("template");
        template.innerHTML = content;
        this.shadowRoot = this.attachShadow({ mode: "open" });
        this.shadowRoot.appendChild(template.content.cloneNode(true));
    }

    connectedCallback() {
        console.log("LoginForm added to page.");
        // this.shadowRoot.addEventListener("click", loginHandler);
        this.shadowRoot.addEventListener("click", sendPostHandler);
    }

    disconnectedCallback() {
        console.log("LoginForm removed from page.");
    }

    attributeChangedCallback(name, oldValue, newValue) {
        console.log(`Attribute ${name} has changed.`);
        console.log(`${oldValue} -> ${newValue}`);
    }
}

function sendPostHandler(event) {
    if (event.target.id != "send-post") {
        return;
    }
    const root = event.target.getRootNode();
    const host = root.host;
    let input = root.getElementById("text-area");
    let contentData = input.value;
    if (contentData.length <= 0) {
        alert("Input cannot be empty.");
        return;
    }
    let cat = [];
    let cat1 = root.getElementById("category1");
    if (cat1.checked) {
        cat.push(cat1.value);
    }
    let cat2 = root.getElementById("category2");
    if (cat2.checked) {
        cat.push(cat2.value);
    }
    let cat3 = root.getElementById("category3");
    if (cat3.checked) {
        cat.push(cat3.value);
    }
    if (cat.length <= 0) {
        alert("One or more categories must be selected.");
        return;
    }
    let catData = cat.join(", ");
    let postData = {
        authorID: host.User.id,
        author: host.User.nickname,
        content: contentData,
        categories: catData,
        timestamp: Math.floor(Date.now() / 1000)
    };

    let body = {
        code: codeNewPost,
        data: postData
    };
    host.Socket.send(JSON.stringify(body));
}
