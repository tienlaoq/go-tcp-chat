let model = {
    history: ["hello"]
};

function update(action) {
    switch (action.type) {
        case 'ADD':
            model.history.unshift(action.payload);
            render();
            break;
    }
}

function createElement(tag, props, ...children) {
    const element = document.createElement(tag);
    if (props) {
        Object.keys(props).forEach(key => {
            if (key.startsWith('on')) {
                element.addEventListener(key.substring(2).toLowerCase(), props[key]);
            } else {
                element[key] = props[key];
            }
        });
    }
    children.forEach(child => {
        if (typeof child === 'string') {
            element.appendChild(document.createTextNode(child));
        } else {
            element.appendChild(child);
        }
    });
    return element;
}

function render() {
    const app = document.getElementById('app');
    app.innerHTML = '';

    model.history.forEach(item => {
        const textNode = createElement('div', null, item);
        app.appendChild(textNode);
    });

    const input = createElement('input', {
        placeholder: 'Type text',
        oninput: event => {
            const value = event.target.value;
            update({ type: 'ADD', payload: value });
            event.target.value = '';
        }
    });

    app.appendChild(input);
}

render();
