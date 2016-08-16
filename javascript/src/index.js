
import React from 'react';
import ReactDOM from 'react-dom';
import { AreaChart, Area, CartesianGrid, YAxis, XAxis, Tooltip } from "recharts";
import Format from "date-format-lite"
import '../../assets/stylesheets/index.css'

const SimpleLineChart = React.createClass({
	getInitialState: function() {
		return {
			isVisible: false,
			data: {}
		};
	},

	formatDate: function(data) {
		data.Weights.forEach(function(current, index, weights) {
			data.Weights[index].RecordedAt = new Date(current.RecordedAt).format("UTC:DD-MM-YYY HH:mm A");
		})
		return data;
	},

	componentWillMount: function() {
		if(window["WebSocket"]) {
	        let conn = new WebSocket("ws://localhost:8080/api/getWeight");
	        conn.onmessage = function (evt) {
	            this.setState({
	            	isVisible: true,
	            	data: this.formatDate(JSON.parse(evt.data))
	            });
	            setTimeout(function(){
	            	this.setState({
	            		isVisible: false,
	            		data: {}
	            	});
	            }.bind(this), 30*1000);
	        }.bind(this);
    	}
	},

	render: function() {
		if(!this.state.isVisible){
			return ( <h2>Please step on the weighing machine</h2> )
		} else if(this.state.data.IsError){
			return ( <h2>{this.state.data.ErrorMsg}</h2> )
		} else {
			return (
				<AreaChart width={730} height={250} data={this.state.data.Weights} margin={{ top: 10, right: 30, left: 0, bottom: 0 }}>
				  <defs>
				    <linearGradient id="colorUv" x1="0" y1="0" x2="0" y2="1">
				      <stop offset="5%" stopColor="#8884d8" stopOpacity={0.8}/>
				      <stop offset="95%" stopColor="#8884d8" stopOpacity={0}/>
				    </linearGradient>
				  </defs>
				  <XAxis dataKey="RecordedAt" />
				  <YAxis />
				  <CartesianGrid strokeDasharray="3 3" />
				  <Tooltip />
				  <Area type="monotone" dataKey="Weight" stroke="#8884d8" fillOpacity={1} fill="url(#colorUv)" />
				</AreaChart>
			);

		}
	}
});

ReactDOM.render(<SimpleLineChart />, document.querySelector('#main'));