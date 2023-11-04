# JS Web Components

## Constructor

Each component in this application is constructed with arguments
corresponding to its necessary data, and a reference to the websocket
if required.

## connectedCallback

Upon connection to the DOM, a component adds to its shadow root
its event listeners and appropriate event handlers.

## disconnectedCallback

Upon disconnection from the DOM, a component removes from its
shadow root its event listeners.

## attributeChangedCallback

