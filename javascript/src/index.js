
import React from 'react';
import ReactDOM from 'react-dom';
import { LineChart, Line, CartesianGrid, YAxis, XAxis, Tooltip } from "recharts";

const data = [
	{date:"10-07-2016", weight:50 },
	{date:"11-07-2016", weight:60 },
	{date:"12-07-2016", weight:54 },
	{date:"15-07-2016", weight:55 },
	{date:"18-07-2016", weight:58 },
	{date:"20-07-2016", weight:53 }
]

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
	render() {
		return (
			<LineChart width={600} height={300} data={data} margin={{ top: 20, right: 40, bottom: 5, left: 0 }} >
			  <Line type="monotone" dataKey="weight" stroke="#8884d8" xAxisId="dateAxis" yAxisId="weightAxis" unit="Kg" label={<CustomizedLabel />}/>
			  <CartesianGrid stroke="#ccc" strokeDasharray="5 5" />
			  <XAxis dataKey="date" xAxisId="dateAxis" height={60} tick={<CustomizedXAxisTick/>} />
			  <YAxis yAxisId="weightAxis" domain={['dataMin - 10', 'dataMax + 10']} tick={<CustomizedYAxisTick/>} />
			  <Tooltip />
			</LineChart>
		);
	}
})

ReactDOM.render(<SimpleLineChart />, document.querySelector('#main'));