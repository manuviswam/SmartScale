
import React from 'react';
import ReactDOM from 'react-dom';
import { LineChart, Line, CartesianGrid, YAxis, XAxis, Tooltip } from "recharts";
import Format from "date-format-lite"

const CustomizedLabel = React.createClass({
  render () {
    const {x, y, stroke, payload} = this.props;
		
   	return <text x={x} y={y} dy={-4} fill={stroke} fontSize={10} textAnchor="middle">{payload.value}</text>
  }
});

const CustomizedXAxisTick = React.createClass({
  render () {
    const {x, y, stroke, payload} = this.props;
		
   	return (
    	<g transform={`translate(${x},${y})`}>
        <text x={0} y={0} dy={16} textAnchor="end" fill="#666" transform="rotate(-35)" fontSize="10">{payload.value}</text>
      </g>
    );
  }
});

const CustomizedYAxisTick = React.createClass({
  render () {
    const {x, y, stroke, payload} = this.props;
		
   	return (
    	<g transform={`translate(${x},${y})`}>
        <text x={0} y={0} dy={16} textAnchor="end" fill="#666" fontSize="10">{payload.value}</text>
      </g>
    );
  }
});

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
			return (
				<text>Please step on the weighing machine</text>
				)
		}else {
			return (
				<div>
				<text>Hello {this.state.data.EmpName}. Your current weight is {this.state.data.CurrentWeight}</text>
				<LineChart width={600} height={400} data={this.state.data.Weights} margin={{ top: 20, right: 40, bottom: 5, left: 0 }} >
				  <Line type="monotone" dataKey="Weight" stroke="#8884d8" xAxisId="dateAxis" yAxisId="weightAxis" unit="Kg" label={<CustomizedLabel />}/>
				  <CartesianGrid stroke="#ccc" strokeDasharray="5 5" />
				  <XAxis dataKey="RecordedAt" xAxisId="dateAxis" height={100} tick={<CustomizedXAxisTick/>} />
				  <YAxis yAxisId="weightAxis" domain={['dataMin - 2', 'dataMax + 2']} tick={<CustomizedYAxisTick/>} />
				  <Tooltip />
				</LineChart>
				</div>
			);

		}
	}
});

ReactDOM.render(<SimpleLineChart />, document.querySelector('#main'));