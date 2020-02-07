import ReactDOM from "react-dom";
import React from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:3000/";

const style = (
    <link
        rel="stylesheet"
        href="https://cdn.jsdelivr.net/npm/semantic-ui@2.4.1/dist/semantic.min.css"
    />
);
class ESSearch extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            query: "how",
            hits: 0,
            results: []
        };
    }

    componentDidMount() {
        this.getQueryResults();
        this.interval = setInterval(() => this.getQueryResults(), 100);
    }
    componentWillUnmount() {
        clearInterval(this.interval);
    }
    getQueryResults() {
        let { query } = this.state;
        if (query) {
            axios.get(endpoint + "search?term=" + query).then(res => {
                // console.log(res["data"]["hits"]["hits"]);
                if (res.data) {
                    this.setState({
                        hits: res.data.hits.total,
                        results: res.data.hits.hits
                    });
                } else {
                    this.setState({
                        hits: 0,
                        results: []
                    });
                }
            });
        }
    }
    onSubmit = () => {
        this.getQueryResults();
    };
    onChange = event => {
        this.setState({
            [event.target.name]: event.target.value
        });
    };
    render() {
        return (
            <div>
                <div className="row">
                    <Header className="header" as="h1">
                        SEARCH NOTES
                    </Header>
                </div>
                <div className="row">
                    <Form onSubmit={this.onSubmit}>
                        <Input
                            type="text"
                            name="query"
                            value={this.state.query}
                            onChange={this.onChange}
                            fluid
                            placeholder="Enter Query"
                        />
                        {/* <Button >Create Task</Button> */}
                    </Form>
                </div>
                <div className="hitCount">
                    <h2>{this.state.hits} Hits</h2>
                </div>
                <div class="ui four doubling stackable cards">
                    {this.state.results.map(result => (
                        // <h3>{result["_source"]["title"]}</h3>
                        <Card>
                            <Card.Header>
                                {result["_source"]["title"]}
                            </Card.Header>
                            <Card.Content>
                                {result["_source"]["text"]}
                            </Card.Content>
                        </Card>
                    ))}
                </div>
                {/* <div className="row"> */}
                {/*     <Card.Group>{this.state.results}</Card.Group> */}
                {/* </div> */}
            </div>
        );
    }
}

// ========================================

ReactDOM.render(
    <div>
        {style}
        <ESSearch />
    </div>,
    document.getElementById("root")
);
