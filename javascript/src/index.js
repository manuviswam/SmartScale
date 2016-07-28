
import React from 'react';
import ReactDOM from 'react-dom';
import { LineChart, Line, CartesianGrid, YAxis, XAxis, Tooltip } from "recharts";

const data = [
	{date:"Mon", weight:50 },
	{date:"Tue", weight:60 },
	{date:"Wed", weight:54 },
	{date:"Fri", weight:55 },
	{date:"Sat", weight:58 },
	{date:"Sun", weight:53 }
]

const SimpleLineChart = React.createClass({
	render() {
		return (
			<LineChart width={600} height={300} data={data} margin={{ top: 5, right: 20, bottom: 5, left: 0 }}>
			  <Line type="monotone" dataKey="weight" stroke="#8884d8" />
			  <CartesianGrid stroke="#ccc" strokeDasharray="5 5" />
			  <XAxis dataKey="date" />
			  <YAxis />
			  <Tooltip />
			</LineChart>
		);
	}
})

ReactDOM.render(<SimpleLineChart />, document.querySelector('#main'));