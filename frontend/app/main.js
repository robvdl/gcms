'use strict';

require('./css/app.css');

var React = require('react'),
    Hello = require('./Hello');

React.render(<Hello />, document.getElementById('content'));
