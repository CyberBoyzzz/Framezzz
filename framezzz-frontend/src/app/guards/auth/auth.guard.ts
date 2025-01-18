import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { CookieService } from 'ngx-cookie-service';

export const authGuard: CanActivateFn = () => {
  const router = inject(Router);

  const cookieService = inject(CookieService);

  if (cookieService.check('authToken')) {
    return true;
  }

  router.navigate(['login']);
  return false;
};
