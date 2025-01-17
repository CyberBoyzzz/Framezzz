import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Comic } from '../../interfaces/comic.interface';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ComicsService {
  private readonly httpClient = inject(HttpClient);

  public getComic(id: number): Observable<Comic> {    
    return this.httpClient.get<Comic>(`/api/${id}/info.0.json`);
  }
}
