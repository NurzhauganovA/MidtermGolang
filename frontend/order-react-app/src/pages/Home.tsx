import React, {useEffect} from 'react';

const Home = () => {

    useEffect(() => {
        (
            async () => {
                await fetch('http://localhost:8000/api/categories', {
                    headers: {
                        'Content-Type': 'application/json',
                        'Accept': 'application/json'
                    },
                });
            }
        )();
    })

    return (
        <div>

        </div>
    );
};

export default Home;