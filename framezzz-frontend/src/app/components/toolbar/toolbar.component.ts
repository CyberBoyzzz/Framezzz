import { Component, inject } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatMenuModule } from '@angular/material/menu';
import { Router } from '@angular/router';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-toolbar',
  imports: [MatToolbarModule, MatButtonModule, MatIconModule, MatMenuModule],
  templateUrl: './toolbar.component.html',
  styleUrl: './toolbar.component.sass',
})
export class ToolbarComponent {
  private readonly router = inject(Router);

  private readonly cookieService = inject(CookieService);

  protected openLoginPage() {
    this.router.navigate(['login']);
  }

  protected openRegisterPage() {
    this.router.navigate(['register']);
  }

  protected openHomePage() {
    this.router.navigate(['']);
  }

  protected signOut() {
    this.cookieService.delete('authToken', '/');

    this.router.navigate(['login']);
  }

  protected isLoggedIn() {
    return this.cookieService.check('authToken');
  }
}
