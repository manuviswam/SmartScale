
import React from 'react';
import ReactDOM from 'react-dom';
import { AreaChart, Area, CartesianGrid, YAxis, XAxis, Tooltip } from "recharts";
import Format from "date-format-lite"
import '../../assets/stylesheets/index.css'

const {PropTypes} = React;

var timeout;

const CustomTooltip  = React.createClass({
  propTypes: {
    type: PropTypes.string,
    payload: PropTypes.array,
    label: PropTypes.string,
  },

  render() {
    const { active } = this.props;

    if (active) {
      const { payload, label } = this.props;
      return (
        <div className="custom-tooltip">
          <p className="label">{label}</p>
          <p className="intro">{payload[0].value} Kg </p>
        </div>
      );
    }

    return null;
  }
});

const CustomizedLabel = React.createClass({
  render () {
    const {x, y, stroke, payload} = this.props;
   	return <text x={x} y={y} dy={-4} fill={stroke} fontSize={10} textAnchor="middle">{payload.value[1]}</text>
  }
});

const CustomizedXAxisTick = React.createClass({
  render () {
    const {x, y, stroke, payload} = this.props;
		
   	return (
    	<g transform={`translate(${x},${y})`}>
        <text textAnchor="middle" y={10} fill="#666" fontSize={10}>{payload.value}</text>
      </g>
    );
  }
});

const CustomizedYAxisTick = React.createClass({
  render () {
    const {x, y, stroke, payload} = this.props;
		
   	return (
    	<g transform={`translate(${x},${y})`}>
        <text fill="#666" textAnchor="end" fontSize={10}>{payload.value} Kg</text>
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

	formatData: function(data) {
		if (data.IsError) {
			return data;
		}
		data = this.formatDate(data);
		data = this.strip(data);
		return data;
	},

	formatDate: function(data) {
		data.Weights.forEach(function(current, index, weights) {
			data.Weights[index].RecordedAt = new Date(current.RecordedAt).format("UTC:DD-MM-YYY HH:mm A");
		})
		return data;
	},

	strip: function(data) {
		data.CurrentWeight = parseFloat(parseFloat(data.CurrentWeight).toPrecision(3));
		data.Weights.forEach(function(current, index, weights) {
			data.Weights[index].Weight = parseFloat(parseFloat(data.Weights[index].Weight).toPrecision(3));
		})
		return data;
	},

	componentWillMount: function() {
		if(window["WebSocket"]) {
	        let conn = new WebSocket("ws://localhost:8080/api/getWeight");
	        conn.onmessage = function (evt) {
	            this.setState({
	            	isVisible: true,
	            	data: this.formatData(JSON.parse(evt.data))
	            });
	            if(timeout) {
	            	clearTimeout(timeout);
	            }
	            timeout = setTimeout(function(){
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
			return ( <h2>Please step on the weighing machine and then swipe</h2> )
		} else if(this.state.data.IsError){
			return ( <h2>{this.state.data.ErrorMsg}</h2> )
		} else {
			return (
			 <div id="mainWrapper">
			 	<div id="message" >
			 		<h2>Hello {this.state.data.EmpName} </h2>
			 		<p id="currentWeight">{this.state.data.CurrentWeight}</p>
			 	</div>
				<AreaChart width={730} height={250} data={this.state.data.Weights} margin={{ top: 10, right: 30, left: 0, bottom: 0 }}>
				  <defs>
				    <linearGradient id="colorUv" x1="0" y1="0" x2="0" y2="1">
				      <stop offset="5%" stopColor="#8884d8" stopOpacity={0.8}/>
				      <stop offset="95%" stopColor="#8884d8" stopOpacity={0}/>
				    </linearGradient>
				  </defs>
				  <XAxis dataKey="RecordedAt" tick={<CustomizedXAxisTick/>} />
				  <YAxis type="number" domain={['dataMin - 2', 'dataMax + 2']} tick={<CustomizedYAxisTick/>}  />
				  <CartesianGrid strokeDasharray="3 3" />
				  <Tooltip content={<CustomTooltip /> }/>
				  <Area type="monotone" dataKey="Weight" stroke="#8884d8" fillOpacity={1} fill="url(#colorUv)"  label={<CustomizedLabel />} />
				</AreaChart>
			</div>
			);

		}
	}
});

ReactDOM.render(<SimpleLineChart />, document.querySelector('#main'));