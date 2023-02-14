import React, {useState} from "react";
import { Link } from 'react-router-dom';
import "./PageSelector.css";

const linkStyle = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    textAlign: "center",
    width: "22vw",
    height: "3vh",
    backgroundColor: "#00d",
    textDecoration : "none",
    borderColor: "transparent",
    color: "white",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
    fontSize: "1.4vh",
}

const disabledLinkStyle = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    textAlign: "center",
    width: "22vw",
    height: "3vh",
    backgroundColor: "#bbb",
    textDecoration : "none",
    borderColor: "transparent",
    color: "white",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
    fontSize: "1.4vh",
}


const PageSelector = (props) => {
    const [pageState, setPage] = useState(props.pageState);
    
    return (
        <div className="page-selector">
            {pageState ?
            <Link   
                to={`/feed/`} 
                style={disabledLinkStyle}
                onClick={ (event) => event.preventDefault() }
            >Все шутки</Link>
            :   <Link   to={`/feed/`} 
                        style={linkStyle}
                        onClick={() => (setPage(false))}
            >Все шутки</Link>}
            {pageState ? 
            <Link   
                to={`/subscribes/`}
                style={linkStyle}
                onClick={ () => setPage(true) }
            >Подписки</Link>
            :   <Link   to={`/subscribes/`}
                        style={disabledLinkStyle}
                        onClick={ (event) => event.preventDefault() }
            >Подписки</Link>}
        </div>
    );
}

export default PageSelector;