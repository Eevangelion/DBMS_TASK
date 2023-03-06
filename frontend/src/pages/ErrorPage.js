import '../styles/ErrorPage.css';

const ErrorPage = (props) => {
    
    return (
        <div className="error-page">
            <div className="error-text">
                <h1>Oops!</h1>
                <p>Sorry, an unexpected error has occurred.</p>
                <h2>404 - Not found</h2>
            </div>
        </div>
    );
};

export default ErrorPage;