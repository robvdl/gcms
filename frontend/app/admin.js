'use strict';

require('./css/app.css');
require('./css/admin.css');

var React = require('react'),
    Hello = require('./components/Hello');

React.render(<Hello />, document.getElementById('admin'));
