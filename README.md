![VPDesktop Cover](demo/vpdesktop_cover_bw.png)

## VPDesktop - Native Desktop Client for VpMobil / Stundenplan24


Overview
--------
VPDesktop is a native desktop client for viewing school timetables from VpMobil/Stundenplan24.
Built in Go, it provides a modern, fast, and reliable alternative to the legacy mobile apps.
VPDesktop allows students to view classes for the current or next schoolday with full API integration.

Built With
----------
- Go (Golang)
- jibber_jabber: language detection
- go-i18n: multilingual support (German / English)
- gio: UI framework for native desktop interface
- go-keyring: secure persistent credential storage

Features
--------
- Full API integration (reverse-engineered, fully typed)
- Multilingual interface (German / English)
- Login via UI (no manual configuration)
- View classes for current or next schoolday

Upcoming Features
-----------------
- Persistent and secure authentication
- Enhanced design and UI improvements
- Table views: week view / month view
- Timetable caching for offline access and faster startup

Installation & Building
----------------------
1. Clone the repository:
   ```
   git clone https://github.com/PokeMatPok/VPDesktop
   cd https://github.com/PokeMatPok/VPDesktop
   ```

2. Install dependencies:
   ```
   go install ./...
   ```

4. Build the desktop client:
   ```
   go build -o VPDesktop
   ```

6. Run the app:
   ```
   ./VPDesktop
   ```

Usage
-----
- Launch VPDesktop and log in with your VpMobil/Stundenplan24 credentials.
- Select your class or school.
- View current or next schoolday timetables.
- Future updates will include week/month views and caching.

Contributing
------------
Contributions are welcome! If you want to add features, fix bugs, or improve the UI:
- Fork the repository
- Create a new branch
- Make your changes
- Submit a pull request

License
-------
This project is provided as-is for educational and personal use. Please respect the
original VpMobil/Stundenplan24 service terms. Do not redistribute any private credentials
or sensitive school data.

Contact
-------
For questions, suggestions, or bug reports, please open an issue on the GitHub repository
or contact the author directly.
