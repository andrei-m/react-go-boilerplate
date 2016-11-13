class Hello extends React.Component {
	render() {
		return (
			<div>{this.props.message}</div>
		);
	}
}

class App extends React.Component {
	constructor(props) {
		super(props);
		this.state = {message: ""};

		fetch('/hello').then(function(resp) {
			return resp.json();
		}).then((function(j) {
			this.setState({message: j.message});
		}).bind(this));
	}

	render() {
		return (
			<Hello message={this.state.message} />
		)
	}
}

ReactDOM.render(
	<App />,
	document.getElementById('container')
);
