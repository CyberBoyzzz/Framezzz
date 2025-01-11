import { Component } from '@angular/core';
import { ComicsCardComponent } from '../../components/comics-card/comics-card.component';
import { ToolbarComponent } from '../../components/toolbar/toolbar.component';

@Component({
  selector: 'app-home',
  imports: [ComicsCardComponent, ToolbarComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.sass'
})
export class HomeComponent {

}
