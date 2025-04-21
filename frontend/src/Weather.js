import React, { useState, useEffect } from 'react';

function Weather() {
    const [city, setCity] = useState("");
    const [weather, setWeather] = useState(null);
    const [error, setError] = useState(null);

    const fetchWeather = async () => {
        try {
          const response = await fetch(`/weather?city=${city}`);
          if (!response.ok) {
            throw new Error("Ошибка при получении данных о погоде");
          }
          const data = await response.json();
          setWeather(data);
          setError(null);
        } catch (err) {
          setError(err.message);
          setWeather(null);
        }
      };
    
      return (
        <div>
          <h1>Погода</h1>
          <input
            type="text"
            placeholder="Введите город"
            value={city}
            onChange={(e) => setCity(e.target.value)}
          />
          <button onClick={fetchWeather}>Получить погоду</button>
          {error && <p style={{ color: "red" }}>{error}</p>}
          {weather && (
            <div>
              <h2>{weather.city}</h2>
              <p>Дата: {weather.datetime}</p>
              <p>Температура: {weather.temp}°C</p>
            </div>
          )}
        </div>
      );
}