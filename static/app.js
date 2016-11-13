class HelloWorld extends React.Component {
	render() {
		return (
			<div>Hello, World!</div>
		);
	}
}

ReactDOM.render(
	<HelloWorld />,
	document.getElementById('app')
);
