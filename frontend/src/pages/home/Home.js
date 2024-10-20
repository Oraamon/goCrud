import React from "react";
import "./WelcomeScreen.css";

const WelcomeScreen = () => {
    return (
        <div className="welcome-container">
            <div className="welcome-content">
                <div className="image-container">
                    <img
                        src=""
                        alt="R Graphic"
                        className="welcome-image"
                    />
                </div>
                <div className="text-container">
                    <div className="title-container">
                        <div className="welcome-title">The <br/> Ramori</div>
                        <h1 className="welcome-message">Welcome!!</h1>
                    </div>
                    <button className="go-button">GOOO!</button>
                    <p className="welcome-description">
                        Discover exclusive products thoughtfully designed to inspire and
                        enhance your everyday life. Explore our collections and find items
                        that perfectly match your style and needs.
                    </p>
                </div>
            </div>
        </div>
    );
};

export default WelcomeScreen;
