import ReactDOM from "react-dom";
import React from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:3000/";

const style = (
    <link
        rel="stylesheet"
        href="//cdn.jsdelivr.net/npm/semantic-ui@2.4.2/dist/semantic.min.css"
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
                        Search My Notes
                    </Header>
                </div>
                <div className="row">
                    <Form>
                        <Input
                            type="text"
                            name="query"
                            value={this.state.query}
                            onChange={this.onChange}
                            icon="search"
                            fluid
                            placeholder="Enter Query"
                        />
                    </Form>
                </div>
                <div className="hitCount">
                    <Header className="header" as="h2">
                        {this.state.hits} Hits
                    </Header>
                </div>
                <div class="ui four doubling stackable cards">
                    {this.state.results.map(result => (
                        <Card>
                            <Card.Header>
                                <strong>{result._source.title}</strong>
                            </Card.Header>
                            <Card.Content>
                                <div
                                    dangerouslySetInnerHTML={{
                                        __html: result.highlight.text.join(
                                            "<br><br>"
                                        )
                                    }}
                                />
                            </Card.Content>
                        </Card>
                    ))}
                </div>
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
