import logo from "../../styles/img/logo_test.png";
import TopPanelButtons from "./TopPanelButtons";
import "./TopPanel.css";

function TopPanel() {
    return (
    <div className="top-panel">
        <a className="main-page-redirect" href="/">
            <img className="main-page-redirect-image" src={logo} alt=":("/>
        </a>
        
        <div className="search-panel">
            <form action="/search/" autoComplete="off" className="form-search" method="get" role="search">
                <input type="search" className="search" placeholder="Поиск" />
            </form>
        </div>
        
        <TopPanelButtons isAuth={true}/>
    </div>)
}
export default TopPanel;