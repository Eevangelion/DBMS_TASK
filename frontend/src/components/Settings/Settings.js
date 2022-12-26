import { useNavigate, useLocation } from "react-router-dom"
import { useGetUserByNameQuery } from "../../services/User";
import {Link} from 'semantic-ui-react';
import Profile from "../Profile/Profile";
import './CreateJoke.css';

const linkStyle = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    textAlign: "center",
    width: "380px",
    height: "30px",
    borderRadius: "45px",
    backgroundColor: "#00d",
    textDecoration : "none",
    borderColor: "transparent",
    color: "white",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
}


const Settings = () => {
    const navigate = useNavigate();
    const location = useLocation();

    const userName = localStorage.getItem('userName');

    const user = useGetUserByNameQuery(userName);

    return (
        <div className="modal-window">
            <div className="buttons">
                <Link to={`settings/developer`} 
                      style={linkStyle}
                      state={{ backgroundLocation: location }}>
                    <strong>Настройки разработчика</strong>
                </Link>
                <button className="back-button" onClick={navigate(-1)}>
                    Назад
                </button>
            </div>
            <Profile user={user} />
        </div>
    )
}

export default Settings;