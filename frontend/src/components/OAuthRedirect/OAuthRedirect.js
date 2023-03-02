import { useNavigate } from "react-router-dom";
import { useGetGitQuery } from "../../services/auth"
import { getCode } from "../../store/actions/auth";
import LoadingModal from "../LoadingModal/LoadingModal";

const OAuthRedirect = () => {
    const navigate = useNavigate();
    const code = getCode();
    const {
        data: tokens,
        isLoading: loadingCode
    } = useGetGitQuery(code);
    if (loadingCode) {
        return <LoadingModal />;
    }
    const accessToken = tokens.jwt_token;
    const refreshToken = tokens.refresh_token;
    const base64Url = accessToken.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(window.atob(base64).split('').map((c) => {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));
    const user = JSON.parse(jsonPayload);
    localStorage.setItem("userID", user.user_id);
    localStorage.setItem("userName", user.username);
    localStorage.setItem("userRole", user.role);
    localStorage.setItem("access_token", accessToken);
    localStorage.setItem("token_exp_time", user.exp);
    localStorage.setItem("refresh_token", refreshToken);
    console.log(user);
    navigate(`/feed/`);
    return;
}

export default OAuthRedirect;