            <table>
                <thead>
                    <tr>
                        <th>POS</th>
                        <th>NO</th>
                        <th>DRIVER</th>
                        <th>CAR</th>
                        <th>LAPS</th>
                        <th>TIME/RETIRED</th>
                        <th>PTS</th>
                    </tr>
                </thead>
                <tbody>
                    
                    {{ range .RaceData.Results }}
                    <tr>
                        <td>{{ .Position }}</td>
                        <td>{{ .Driver.Number }}</td>
                        <td>{{ .Driver.FirstName }} {{ .Driver.FamilyName }}</td>
                        <td>{{ .Constructor.Name }}</td>
                        <td>{{ .Laps }}</td>
                        <td>{{ if ne .Time.Time "" }}
                                {{ .Time.Time }}
                            {{else }}
                                {{$firstChar := index .Status 0}}
                                {{ if ne $firstChar 43}}DNF - {{end}}{{ .Status }}
                            {{ end }}</td>
                        <td>{{ .Points }}</td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
            
